package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"proyecto1/cors"

	"os/exec"

	"github.com/gorilla/mux"
)

func RamController(w http.ResponseWriter, r *http.Request) {
	// Ejecución del comando
	cmd := exec.Command("cat", "/proc/ram_so1_1s2024")
	salida, err := cmd.Output()
	if err != nil {
		// Manejar el error
		fmt.Fprintf(w, "Error al obtener la información de la RAM: %v", err)
		return
	}

	// Decodificación de la salida JSON
	var data map[string]interface{}
	err = json.Unmarshal(salida, &data)
	if err != nil {
		// Manejar el error
		fmt.Fprintf(w, "Error al decodificar la información de la RAM: %v", err)
		return
	}

	// Conversión a JSON y envío de la respuesta
	jsonData, err := json.Marshal(data)
	if err != nil {
		// Manejar el error
		fmt.Fprintf(w, "Error al convertir la información a JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func CpuController(w http.ResponseWriter, r *http.Request) {
	// Ejecución del comando
	cmd := exec.Command("cat", "/proc/cpu_so1_1s2024")
	salida, err := cmd.Output()
	if err != nil {
		// Manejar el error
		fmt.Fprintf(w, "Error al obtener la información de la CPU: %v", err)
		return
	}

	// Decodificación de la salida JSON
	var data map[string]interface{}
	err = json.Unmarshal(salida, &data)
	if err != nil {
		// Manejar el error
		fmt.Fprintf(w, "Error al decodificar la información de la CPU: %v", err)
		return
	}

	// Conversión a JSON y envío de la respuesta
	jsonData, err := json.Marshal(data)
	if err != nil {
		// Manejar el error
		fmt.Fprintf(w, "Error al convertir la información a JSON: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func InitializeRoutes(router *mux.Router) {
	router.HandleFunc("/cpu", CpuController).Methods("GET")
	router.HandleFunc("/ram", RamController).Methods("GET")

	// Aplica la configuración de CORS a todas las rutas
	router.Use(cors.CorsHandler())
}
