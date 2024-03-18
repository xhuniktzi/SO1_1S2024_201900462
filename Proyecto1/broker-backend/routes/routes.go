package routes

import (
	"broker-backend/config"
	"broker-backend/cors"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"os/exec"

	"github.com/gorilla/mux"
)

var process *exec.Cmd

func StartSignal(w http.ResponseWriter, r *http.Request) {
	// Crear un nuevo proceso con un comando de espera
	cmd := exec.Command("sleep", "infinity")
	err := cmd.Start()
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Error al iniciar el proceso", http.StatusInternalServerError)
		return
	}

	// Obtener el comando con PID
	process = cmd

	fmt.Fprintf(w, "%d", process.Process.Pid)
}

func StopSignal(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGSTOP al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-SIGSTOP", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al detener el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Proceso con PID %d detenido", pid)
}

func ResumeSignal(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGCONT al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-SIGCONT", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al reanudar el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Proceso con PID %d reanudado", pid)
}

func KillSignal(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	// Enviar señal SIGCONT al proceso con el PID proporcionado
	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al intentar terminar el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Proceso con PID %d ha terminado", pid)
}

func DataController(w http.ResponseWriter, r *http.Request) {
	// Configuración de la conexión a la base de datos
	db := config.GetDb()

	// Consulta para obtener los últimos 20 registros de la tabla CPU
	rowsCPU, err := db.Query("SELECT * FROM (SELECT * FROM cpu ORDER BY id DESC LIMIT 20) sub ORDER BY id ASC")
	if err != nil {
		// Manejar el error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rowsCPU.Close()

	// Consulta para obtener los últimos 20 registros de la tabla RAM
	rowsRAM, err := db.Query("SELECT * FROM (SELECT * FROM ram ORDER BY id DESC LIMIT 20) sub ORDER BY id ASC")
	if err != nil {
		// Manejar el error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rowsRAM.Close()

	// Variables para almacenar los datos
	var cpuData []map[string]interface{}
	var ramData []map[string]interface{}

	// Recuperar datos de la tabla CPU
	for rowsCPU.Next() {
		var id int
		var free, used float64
		err := rowsCPU.Scan(&id, &free, &used)
		if err != nil {
			// Manejar el error
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cpuData = append(cpuData, map[string]interface{}{
			"id":   id,
			"free": free,
			"used": used,
		})
	}

	// Recuperar datos de la tabla RAM
	for rowsRAM.Next() {
		var id int
		var free, used float64
		err := rowsRAM.Scan(&id, &free, &used)
		if err != nil {
			// Manejar el error
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ramData = append(ramData, map[string]interface{}{
			"id":   id,
			"free": free,
			"used": used,
		})
	}

	// Combinar los datos de CPU y RAM en una sola estructura
	responseData := map[string]interface{}{
		"cpu": cpuData,
		"ram": ramData,
	}

	// Convertir la estructura de datos a JSON y enviar como respuesta
	jsonResponse, err := json.Marshal(responseData)
	if err != nil {
		// Manejar el error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer encabezado y escribir respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

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
	router.HandleFunc("/data", DataController).Methods("GET")
	router.HandleFunc("/start", StartSignal).Methods("GET")
	router.HandleFunc("/stop", StopSignal).Methods("GET")
	router.HandleFunc("/resume", ResumeSignal).Methods("GET")
	router.HandleFunc("/kill", KillSignal).Methods("GET")

	// Aplica la configuración de CORS a todas las rutas
	router.Use(cors.CorsHandler())
}
