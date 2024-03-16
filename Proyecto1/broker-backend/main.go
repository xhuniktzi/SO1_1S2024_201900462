package main

import (
	"broker-backend/config"
	"broker-backend/routes"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	config.OpenDB()
	// Función periódica para insertar datos
	go func() {
		ticker := time.NewTicker(10 * time.Second) // Ajusta la duración según necesites
		for range ticker.C {
			insertarDatos()
		}
	}()

	router := mux.NewRouter()
	routes.InitializeRoutes(router)

	log.Println("Servidor corriendo en el puerto 8080")
	http.ListenAndServe(":8080", router)
}

func insertarDatos() {
	// Aquí va tu lógica para insertar datos en la base de datos
	db := config.GetDb()
	fmt.Println("Insertando datos...")
	// Por ejemplo: db.Exec("INSERT INTO tabla (columna) VALUES (valor)")
	cmd1 := exec.Command("cat", "/proc/ram_so1_1s2024")
	salida1, err1 := cmd1.Output()

	if err1 != nil {
		fmt.Println("Error al obtener la información de la RAM:", err1)
		return
	}

	var data1 map[string]interface{}
	err2 := json.Unmarshal(salida1, &data1)
	if err2 != nil {
		fmt.Println("Error al decodificar la información de la RAM:", err2)
		return
	}

	ram_used := data1["UsagePercent"]
	ram_free := 100 - ram_used.(float64)

	fmt.Println("RAM", ram_used, ram_free)
	_, errdb1 := db.Exec("INSERT INTO ram (used, free) VALUES (?, ?)", ram_used, ram_free)

	if errdb1 != nil {
		fmt.Println("Error al insertar datos en la base de datos:", errdb1)
		return
	}

	cmd2 := exec.Command("cat", "/proc/cpu_so1_1s2024")
	salida2, err3 := cmd2.Output()

	if err3 != nil {
		fmt.Println("Error al obtener la información de la CPU:", err3)
		return
	}

	var data2 map[string]interface{}
	err4 := json.Unmarshal(salida2, &data2)
	if err4 != nil {
		fmt.Println("Error al decodificar la información de la CPU:", err4)
		return
	}

	cpu_used := data2["Total_CPU_Time"]

	cpu_cores := 2
	cpu_usage := cpu_used.(float64) / float64(cpu_cores)

	if cpu_usage > 1 {
		cpu_usage = 1
	} else {
		cpu_usage = cpu_usage * 100
	}

	cpu_free := 100 - cpu_usage

	fmt.Println("CPU", cpu_usage, cpu_free)
	_, errdb2 := db.Exec("INSERT INTO cpu (used, free) VALUES (?, ?)", cpu_usage, cpu_free)
	if errdb2 != nil {
		fmt.Println("Error al insertar datos en la base de datos:", errdb2)
		return
	}

	fmt.Println("Datos insertados")
}
