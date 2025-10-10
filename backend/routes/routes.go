package routes

import (
	"github.com/gofiber/fiber/v2"
	"audiovisualpro/backend/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	clientes := api.Group("/clientes")
	{
		clientes.Get("/", handlers.GetClientes)
		clientes.Get("/:id_cliente", handlers.GetCliente)
	}

	auth := api.Group("/auth")
	{
		auth.Post("/login", handlers.Login)
	}
}