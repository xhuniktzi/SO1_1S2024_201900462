// server/main.go
package main

import (
	"context"
	"log"
	"net"

	pb "golang-producer/grpc"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

type server struct {
	pb.UnimplementedGetInfoServer
}

// Configuración del productor de Kafka.
var kafkaWriter = kafka.NewWriter(kafka.WriterConfig{
	Brokers: []string{"kafka:9092"}, // Asegúrate de usar la dirección correcta del broker de Kafka
	Topic:   "vote-topic",
})

// Función para enviar un mensaje a Kafka
func sendMessageToKafka(message []byte) error {
	// Escribe el mensaje en Kafka
	err := kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *server) ReturnInfo(ctx context.Context, in *pb.RequestId) (*pb.ReplyInfo, error) {
	log.Printf("Recibido del cliente gRPC: %v", in)

	// Serializa el mensaje gRPC a bytes
	messageBytes, err := proto.Marshal(in)
	if err != nil {
		log.Printf("Error al serializar el mensaje gRPC: %v", err)
		return nil, err
	}

	// Envía el mensaje serializado a Kafka
	err = sendMessageToKafka(messageBytes)
	if err != nil {
		log.Printf("Error al enviar mensaje a Kafka: %v", err)
		return nil, err
	}

	return &pb.ReplyInfo{Info: "Mensaje recibido y enviado a Kafka"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGetInfoServer(s, &server{})

	log.Printf("Servidor gRPC escuchando en %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
