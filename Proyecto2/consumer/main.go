// consumer/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "consumer/grpc"

	"github.com/go-redis/redis"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Configuración de Kafka
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"kafka:9092"},
		Topic:       "vote-topic",
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
		MaxAttempts: 5,    // Número máximo de reintentos
	})
	defer r.Close()

	// Configuración de Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Configuración de MongoDB
	clientOptions := options.Client().
		ApplyURI("mongodb://mongodb:27017").
		SetMaxPoolSize(200).                  // Mantiene el tamaño máximo del pool
		SetMaxConnIdleTime(60 * time.Second). // Mantiene las conexiones inactivas durante 1 minuto
		SetConnectTimeout(30 * time.Second).  // Incrementa el tiempo de espera para conectar
		SetSocketTimeout(60 * time.Second)    // Mantiene el tiempo de espera para las operaciones en el socket

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error al conectar a MongoDB: %v", err)
	}
	defer mongoClient.Disconnect(ctx)

	// Verifica la conexión
	ctxPing, cancelPing := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelPing()
	if err := mongoClient.Ping(ctxPing, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	db := mongoClient.Database("votes")
	collection := db.Collection("dataset-votes")

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {

			log.Printf("Error al leer un mensaje de Kafka: %v", err)
			continue
		}

		// Deserializa el mensaje de Kafka
		var requestId pb.RequestId
		err = proto.Unmarshal(m.Value, &requestId)
		if err != nil {
			log.Printf("Error al deserializar el mensaje de Kafka: %v", err)
			continue
		}

		// Guarda el mensaje en Redis
		year, err := strconv.Atoi(requestId.Year)
		if err != nil {
			log.Printf("Error al convertir el año a entero: %v", err)
			continue
		}
		err = redisClient.HSet("votes", fmt.Sprintf("%d", year), fmt.Sprintf("%s|%s|%s|%s", requestId.Album, requestId.Artist, requestId.Ranked, requestId.Year)).Err()
		if err != nil {
			log.Printf("Error al guardar en Redis: %v", err)
		}

		// Guarda el mensaje en MongoDB
		_, err = collection.InsertOne(ctx, bson.M{
			"album":  requestId.Album,
			"year":   requestId.Year,
			"artist": requestId.Artist,
			"ranked": requestId.Ranked,
		})
		if err != nil {
			log.Printf("Error al guardar en MongoDB: %v", err)
			continue
		}

		// Store the pointer to requestId in a temporary variable before printing
		tempRequestID := &requestId
		fmt.Printf("Mensaje recibido y guardado: %+v\n", tempRequestID)
	}
}
