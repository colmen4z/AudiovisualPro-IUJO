package main

import (
	"log"

	"github.com/joho/godotenv"
	"audiovisualpro/backend/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Hubo un error al cargar el archivo .env.\nError: %v\n", err)
	}

	database.ConnectDB()
}