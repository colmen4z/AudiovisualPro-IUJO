package database

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"audiovisualpro/backend/models"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Hubo un error al intentar conectarse a la base de datos.\nError: %v\n", err)
	}

	fmt.Printf("Conectado a la base de datos exitosamente.\n")

	DB.AutoMigrate(
		//CATALOGOS
		&models.RolPersonal{},
		&models.EstadoProyecto{},
		&models.EstadoEntregable{},
		&models.EstadoPago{},
		&models.TipoProyecto{},
		&models.TipoRecurso{},
		&models.CategoriaGasto{},
		//ENTIDADES PRINCIPALES
		&models.Cliente{},
		&models.Locacion{},
		&models.Proyecto{},
		&models.Personal{},
		&models.Contrato{},
		&models.Entregable{},
		&models.RecursoTecnico{},
		//RELACIONES
		&models.AsignacionPersonal{},
		&models.UsoRecurso{},
		//CONTABLE
		&models.Factura{},
		&models.Gasto{},
		&models.PagoPersonal{},
		//GESTOR
		&models.Gestor{},
	)

	createDefaultUser()
}

func createDefaultUser() {
	defaultUser := os.Getenv("DEFAULT_USER")
	defaultPassword := os.Getenv("DEFAULT_PASSWORD")

	if defaultUser == "" || defaultPassword == "" {
		log.Fatalf("Las variables de entorno 'DEFAULT_USER' y 'DEFAULT_PASSWORD' no estan configuradas.")
		return
	}

	var gestor models.Gestor
	result := DB.Where("usuario = ?", defaultUser).First(&gestor)

	if result.Error == gorm.ErrRecordNotFound {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Error al hashear la contraseña: %v\n", err)
		}

		newGestor := models.Gestor{
			Nombre: "John Doe",
			Usuario: defaultUser,
			Contrasena: string(hashedPassword),
			Email: "admin@audiovisualpro.com",
			Telefono: "0412-1234567",
		}

		if err := DB.Create(&newGestor).Error; err != nil {
			log.Fatalf("Error al crear usuario predeterminado: %v\n", err)
		}

		fmt.Printf("Gestor '%v' creado con exito", defaultUser)
	} else if result.Error != nil {
		log.Printf("Error al verificar el gestor: %v\n", result.Error)
	} else {
		fmt.Printf("Gestor '%v' ya existe", defaultUser)
	}
}