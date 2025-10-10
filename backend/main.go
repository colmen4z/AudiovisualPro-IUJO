package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"audiovisualpro/backend/database"
	"audiovisualpro/backend/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Hubo un error al cargar el archivo .env.\nError: %v\n", err)
	}

	if os.Getenv("JWT_SECRET") == "" {
		log.Fatalf("No se encuentra la variable de entorno 'JWT_SECRET'")
	}

	database.ConnectDB()

	app := fiber.New()

	app.Use(func (c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.Next()
	})

	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen(":" + port))
}