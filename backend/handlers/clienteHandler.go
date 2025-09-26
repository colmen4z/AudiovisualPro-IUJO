package handlers

import (
	"github.com/gofiber/fiber/v2"
	"audiovisualpro/backend/database"
	"audiovisualpro/backend/models"
)

func GetCliente(c *fiber.Ctx) error {
	id := c.Params("id_cliente")
	var cliente models.Cliente

	result := database.DB.First(&cliente, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"message": "Cliente no encontrado",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"message": "Cliente encontrado",
		"data": cliente,
	})
}

func GetClientes(c *fiber.Ctx) error {
	var clientes []models.Cliente

	result := database.DB.Find(&clientes)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"message": "Clientes no encontrados",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"message": "Clientes encontrados",
		"data": clientes,
		"count": len(clientes),
	})
}

func CreateCliente(c *fiber.Ctx) error {
	var cliente models.Cliente
	if err := c.BodyParser(&cliente); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": true,
			"message": "No se puede analizar el JSON",
		})
	}

	database.DB.Create(&cliente)
	return c.JSON(cliente)
}

func DeleteCliente(c *fiber.Ctx) error {
	id := c.Params("id_cliente")
	var cliente models.Cliente

	result := database.DB.First(&cliente, id)
	if result.Error != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": true,
			"message": "Cliente no encontrado",
		})
	}

	result = database.DB.Delete(&cliente)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": true,
			"message": "No se ha podido eliminar el cliente",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"message": "Cliente eliminado con exito",
		"data": cliente,
	})
}