package config

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// Configuración de la conexión a la base de datos
var db *sql.DB
var err error

func OpenDB() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "rootpassword",
		Net:                  "tcp",
		Addr:                 "mysql:3306",
		DBName:               "monitor",
		AllowNativePasswords: true,
	}
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

}

func GetDb() *sql.DB {
	return db
}
