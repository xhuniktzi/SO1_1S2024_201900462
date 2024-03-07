package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Student struct {
	Carnet string `json:"carnet"`
	Nombre string `json:"nombre"`
}

// Función para manejar el endpoint /data
func dataHandler(w http.ResponseWriter, r *http.Request) {

	student := Student{
		Carnet: "201900462",
		Nombre: "Xhunik Miguel",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

// Middleware para manejar CORS
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Pre-flight request
		if r.Method == "OPTIONS" {
			return
		}

		next(w, r)
	}
}

func main() {
	http.HandleFunc("/data", enableCORS(dataHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
