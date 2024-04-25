package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	pb "consumer/grpc"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/proto"
)

type Data struct {
	Album  string `json:"album"`
	Artist string `json:"artist"`
	Ranked int    `json:"ranked"`
	Year   int    `json:"year"`
}

func main() {
	// Configuración común
	r := setupReader("vote-topic")
	defer r.Close()

	r2 := setupReader("vote-topic2")
	defer r2.Close()

	rdb, mongoClient, logsCollection := setupDatabases()
	defer mongoClient.Disconnect(context.Background())

	// Proceso de lectura y manejo de mensajes
	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			logMessage(logsCollection, "Error al leer un mensaje de Kafka en vote-topic", err)
			// Intenta leer del segundo tópico si falla el primero
			m, err = r2.ReadMessage(context.Background())
			if err != nil {
				logMessage(logsCollection, "Error al leer un mensaje de Kafka en vote-topic2", err)
				continue
			}
			handleJSONMessage(m, rdb, logsCollection)
		} else {
			handleGRPCMessage(m, rdb, logsCollection)
		}
	}
}

func setupReader(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"kafka:9092"},
		Topic:       topic,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		MaxAttempts: 5,    // Número máximo de reintentos
	})
}

func setupDatabases() (*redis.Client, *mongo.Client, *mongo.Collection) {
	// Configuración de Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	// Configuración de MongoDB para logs
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017").SetMaxPoolSize(200)
	mongoClient, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}

	logsCollection := mongoClient.Database("logs").Collection("vote-logs")
	return rdb, mongoClient, logsCollection
}

func handleGRPCMessage(m kafka.Message, rdb *redis.Client, logsCollection *mongo.Collection) {
	var requestId pb.RequestId
	if err := proto.Unmarshal(m.Value, &requestId); err != nil {
		logMessage(logsCollection, "Error al deserializar el mensaje de Kafka (gRPC)", err)
		return
	}
	ranked, _ := strconv.Atoi(requestId.Ranked) // Convert string to int
	year, _ := strconv.Atoi(requestId.Year)     // Convert string to int
	saveMessageData(rdb, logsCollection, requestId.Album, requestId.Artist, ranked, year)
}

func handleJSONMessage(m kafka.Message, rdb *redis.Client, logsCollection *mongo.Collection) {
	cleanedMessage := strings.Trim(string(m.Value), "\"")
	cleanedMessage = strings.ReplaceAll(cleanedMessage, `\"`, `"`)

	log.Printf("Cleaned JSON message: %s", cleanedMessage)

	var data Data
	if err := json.Unmarshal([]byte(cleanedMessage), &data); err != nil {
		logMessage(logsCollection, "Error al deserializar el mensaje de Kafka (JSON)", err)
		return
	}
	saveMessageData(rdb, logsCollection, data.Album, data.Artist, data.Ranked, data.Year)
}

func saveMessageData(rdb *redis.Client, logsCollection *mongo.Collection, album, artist string, ranked, year int) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		logMessage(logsCollection, "Error al generar un UUID", err)
		return
	}

	data := fmt.Sprintf(`uuid: "%s", album: "%s", year: "%d", artist: "%s", ranked: "%d"`, uuid, album, year, artist, ranked)
	_, err = rdb.LPush(context.Background(), "votes_list", data).Result()
	if err != nil {
		logMessage(logsCollection, "Error al guardar en Redis", err)
		return
	}

	logMessage(logsCollection, "Mensaje recibido y guardado", nil)
}

func logMessage(collection *mongo.Collection, message string, err error) {
	logEntry := bson.M{
		"timestamp": time.Now(),
		"message":   message,
		"error":     fmt.Sprint(err),
	}

	if _, err := collection.InsertOne(context.Background(), logEntry); err != nil {
		log.Printf("Error al guardar log en MongoDB: %v", err)
	}

	log.Printf("%s: %v", message, fmt.Sprint(err))
}
