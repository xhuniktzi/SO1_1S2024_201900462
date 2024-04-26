// client/main.go
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	pb "golang-client/grpc"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Data struct {
	Album string `json:"album"`
	Year  string `json:"year"`
	Name  string `json:"name"`
	Rank  string `json:"rank"`
}

// Handler que recibe las peticiones REST y las convierte a llamadas gRPC.
func insertDataHandler(w http.ResponseWriter, r *http.Request) {
	var data Data
	// Decodifica el JSON del cuerpo de la petición.
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Establece la conexión con el servidor gRPC.
	conn, err := grpc.Dial("golang-producer:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("No se pudo conectar al servidor gRPC: %v", err)
	}
	defer conn.Close()

	client := pb.NewGetInfoClient(conn)

	// Envía la petición al servidor gRPC y recibe la respuesta.
	res, err := client.ReturnInfo(context.Background(), &pb.RequestId{
		Album:  data.Album,
		Year:   data.Year,
		Artist: data.Name,
		Ranked: data.Rank,
	})
	if err != nil {
		log.Fatalf("Error al llamar al servicio gRPC: %v", err)
	}

	// Envía la respuesta del servidor gRPC al cliente REST.
	json.NewEncoder(w).Encode(res)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/insert", insertDataHandler).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8080", // Ajusta la dirección y puerto según necesidad
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
