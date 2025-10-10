package handlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"audiovisualpro/backend/database"
	"audiovisualpro/backend/models"
	"audiovisualpro/backend/utils"
)

func Login(c *fiber.Ctx) error {
	var loginRequest models.LoginRequest

	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "Solicitud invalida.",
		})
	}

	loginRequest.Usuario = strings.TrimSpace(loginRequest.Usuario)
	loginRequest.Contrasena = strings.TrimSpace(loginRequest.Contrasena)

	if loginRequest.Usuario == "" || loginRequest.Contrasena == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"message": "El usuario y la contraseña son campos obligatorios.",
		})
	}

	var gestor models.Gestor
	result := database.DB.Where("usuario = ?", loginRequest.Usuario).First(&gestor)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"message": "Datos invalidos. Usuario o contraseña incorrectos.",
			})
		}
		fmt.Printf("Error al consultar la base de datos: %v\n", result.Error)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": "Error interno del servidor.",
		})
	}

	err := bcrypt.CompareHashAndPassword([]byte(gestor.Contrasena), []byte(loginRequest.Contrasena))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"message": "Datos invalidos. Usuario o contraseña incorrectos.",
		})
	}

	token, err := utils.GenerateJWT(gestor.ID, gestor.Usuario)
	if err != nil {
		fmt.Printf("Error al generar el JWT: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"message": "Error al generar el JSON Web Token.",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"message": "Inicio de sesion exitoso.",
		"token": token,
	})
}