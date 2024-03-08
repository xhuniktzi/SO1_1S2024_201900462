package main

import (
	"log"
	"net/http"
	"proyecto1/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Servidor corriendo en el puerto 8080")
	http.ListenAndServe(":8080", router)
}
