// consumer/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "consumer/grpc"

	"github.com/google/uuid"
	"github.com/nitishm/go-rejson/v4"
	"github.com/nitishm/go-rejson/v4/clients"
	"github.com/redis/go-redis/v9"
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
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Cambiar si es necesario
		Password: "",           // no password set
		DB:       0,            // use default DB
	})
	rh := rejson.NewReJSONHandler()
	rh.SetGoRedisClientWithContext(context.Background(), clients.GoRedisClientConn(rdb))

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

		// // Crea el objeto requestId a partir del mensaje deserializado
		// requestIdObj := bson.M{
		// 	"album":  requestId.Album,
		// 	"year":   requestId.Year,
		// 	"artist": requestId.Artist,
		// 	"ranked": requestId.Ranked,
		// }

		// // Guarda el nuevo objeto requestId en el array JSON en Redis usando JSON.ARRAPPEND
		// _, err = rh.JSONArrAppend("votes_array", ".", requestIdObj)
		// if err != nil {
		// 	log.Printf("Error al guardar en Redis: %v", err)
		// 	continue
		// }

		uuid, err := uuid.NewRandom()
		if err != nil {
			log.Printf("Error al generar un UUID: %v", err)
			continue
		}

		data := fmt.Sprintf(`uuid: "%s", album: "%s", year: "%s", artist: "%s", ranked: %s`, uuid, requestId.Album, requestId.Year, requestId.Artist, requestId.Ranked)

		_, err = rdb.LPush(context.Background(), "votes_list", data).Result()

		if err != nil {
			log.Printf("Error al guardar en Redis: %v", err)
			continue
		}

		// Crea un nuevo contexto para la operación de inserción en MongoDB
		insertCtx, insertCancel := context.WithTimeout(context.Background(), 5*time.Second)
		// defer insertCancel()

		_, err = collection.InsertOne(insertCtx, bson.M{
			"album":  requestId.Album,
			"year":   requestId.Year,
			"artist": requestId.Artist,
			"ranked": requestId.Ranked,
		})
		if err != nil {
			log.Printf("Error al guardar en MongoDB: %v", err)
			continue
		}

		// Asegúrate de cancelar el contexto al final de la iteración
		insertCancel()
		fmt.Printf("Mensaje recibido y guardado: %+v\n", &requestId)
	}
}
