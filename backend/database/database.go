package database

import (
	"fmt"
	"log"
	"os"

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
}