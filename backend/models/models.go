package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CATÁLOGOS

type RolPersonal struct {
	IDRol     uint   `json:"id_rol" gorm:"primaryKey;autoIncrement"`
	NombreRol string `json:"nombre_rol" gorm:"size:50;unique;not null"`
}

type EstadoProyecto struct {
	IDEstadoProyecto uint   `json:"id_estado_proyecto" gorm:"primaryKey;autoIncrement"`
	NombreEstado     string `json:"nombre_estado" gorm:"size:50;unique;not null"`
}

type EstadoEntregable struct {
	IDEstadoEntregable uint   `json:"id_estado_entregable" gorm:"primaryKey;autoIncrement"`
	NombreEstado       string `json:"nombre_estado" gorm:"size:50;unique;not null"`
}

type EstadoPago struct {
	IDEstadoPago uint   `json:"id_estado_pago" gorm:"primaryKey;autoIncrement"`
	NombreEstado string `json:"nombre_estado" gorm:"size:50;unique;not null"`
}

type TipoProyecto struct {
	IDTipoProyecto uint   `json:"id_tipo_proyecto" gorm:"primaryKey;autoIncrement"`
	NombreTipo     string `json:"nombre_tipo" gorm:"size:50;unique;not null"`
}

type TipoRecurso struct {
	IDTipoRecurso uint   `json:"id_tipo_recurso" gorm:"primaryKey;autoIncrement"`
	NombreTipo    string `json:"nombre_tipo" gorm:"size:50;unique;not null"`
}

type CategoriaGasto struct {
	IDCategoriaGasto uint   `json:"id_categoria_gasto" gorm:"primaryKey;autoIncrement"`
	NombreCategoria  string `json:"nombre_categoria" gorm:"size:50;unique;not null"`
}

// ENTIDADES PRINCIPALES

type Cliente struct {
	IDCliente       uint      `json:"id_cliente" gorm:"primaryKey;autoIncrement"`
	NombreCliente   string    `json:"nombre_cliente" gorm:"size:100;not null"`
	RIFCliente		string	  `json:"rif_cliente" gorm:"size:10;not null"`
	EmailCliente    string    `json:"email_cliente" gorm:"size:100;unique"`
	TelefonoCliente string    `json:"telefono_cliente" gorm:"size:30"`
	Proyectos       []Proyecto `gorm:"foreignKey:ClienteID"`
	Contratos       []Contrato `gorm:"foreignKey:ClienteID"`
	Facturas        []Factura  `gorm:"foreignKey:ClienteID"`
}

type Locacion struct {
	IDLocacion          uint   `json:"id_locacion" gorm:"primaryKey;autoIncrement"`
	NombreLocacion      string `json:"nombre_locacion" gorm:"size:100;not null"`
	Direccion           string `json:"direccion" gorm:"size:200"`
	DescripcionLocacion string `json:"descripcion_locacion" gorm:"size:255"`
	Proyectos           []Proyecto `gorm:"foreignKey:LocacionID"`
	Recursos            []RecursoTecnico `gorm:"foreignKey:LocacionID"`
}

type Proyecto struct {
	IDProyecto       uint       `json:"id_proyecto" gorm:"primaryKey;autoIncrement"`
	NombreProyecto   string     `json:"nombre_proyecto" gorm:"size:100;not null"`
	TipoProyectoID   uint       `json:"tipo_proyecto_id"`
	EstadoProyectoID uint       `json:"estado_proyecto_id"`
	FechaInicio      time.Time  `json:"fecha_inicio" gorm:"not null"`
	FechaFinEstimada *time.Time `json:"fecha_fin_estimada"`
	Presupuesto      float64    `json:"presupuesto" gorm:"not null"`
	ClienteID        *uint      `json:"cliente_id"`
	LocacionID       *uint      `json:"locacion_id"`

	Cliente        *Cliente        `gorm:"foreignKey:ClienteID"`
	Locacion       *Locacion       `gorm:"foreignKey:LocacionID"`
	EstadoProyecto *EstadoProyecto `gorm:"foreignKey:EstadoProyectoID"`
	TipoProyecto   *TipoProyecto   `gorm:"foreignKey:TipoProyectoID"`

	Contratos    []Contrato           `gorm:"foreignKey:ProyectoID"`
	Entregables  []Entregable         `gorm:"foreignKey:ProyectoID"`
	Asignaciones []AsignacionPersonal `gorm:"foreignKey:ProyectoID"`
	UsoRecursos  []UsoRecurso         `gorm:"foreignKey:ProyectoID"`
	Facturas     []Factura            `gorm:"foreignKey:ProyectoID"`
	Gastos       []Gasto              `gorm:"foreignKey:ProyectoID"`
	Pagos        []PagoPersonal       `gorm:"foreignKey:ProyectoID"`
}

type Personal struct {
	IDPersonal     uint    `json:"id_personal" gorm:"primaryKey;autoIncrement"`
	NombrePersonal string  `json:"nombre_personal" gorm:"size:100;not null"`
	CedulaPersonal string  `json:"cedula_personal" gorm:"size:9;not null"`
	RolID          *uint   `json:"rol_id"`
	Salario        float64 `json:"salario" gorm:"not null"`
	EmailPersonal  string  `json:"email_personal" gorm:"size:100;unique"`
	Telefono       string  `json:"telefono" gorm:"size:30"`

	Rol       *RolPersonal         `gorm:"foreignKey:RolID"`
	Asignados []AsignacionPersonal `gorm:"foreignKey:PersonalID"`
	Pagos     []PagoPersonal       `gorm:"foreignKey:PersonalID"`
}

type Contrato struct {
	IDContrato          uint      `json:"id_contrato" gorm:"primaryKey;autoIncrement"`
	ProyectoID          *uint     `json:"proyecto_id"`
	ClienteID           *uint     `json:"cliente_id"`
	FechaFirma          time.Time `json:"fecha_firma" gorm:"not null"`
	MontoContrato       float64   `json:"monto_contrato" gorm:"not null"`
	DescripcionServicios string   `json:"descripcion_servicios" gorm:"size:255"`

	Proyecto *Proyecto `gorm:"foreignKey:ProyectoID"`
	Cliente  *Cliente  `gorm:"foreignKey:ClienteID"`
}

type Entregable struct {
	IDEntregable        uint       `json:"id_entregable" gorm:"primaryKey;autoIncrement"`
	ProyectoID          uint       `json:"proyecto_id" gorm:"not null"`
	Descripcion         string     `json:"descripcion" gorm:"size:255;not null"`
	FechaEntregaEstimada *time.Time `json:"fecha_entrega_estimada"`
	EstadoEntregableID  *uint      `json:"estado_entregable_id"`

	Proyecto        *Proyecto        `gorm:"foreignKey:ProyectoID"`
	EstadoEntregable *EstadoEntregable `gorm:"foreignKey:EstadoEntregableID"`
}

type RecursoTecnico struct {
	IDRecurso     uint    `json:"id_recurso" gorm:"primaryKey;autoIncrement"`
	NombreEquipo  string  `json:"nombre_equipo" gorm:"size:150;not null"`
	TipoRecursoID *uint   `json:"tipo_recurso_id"`
	LocacionID    *uint   `json:"locacion_id"`

	TipoRecurso *TipoRecurso `gorm:"foreignKey:TipoRecursoID"`
	Locacion    *Locacion    `gorm:"foreignKey:LocacionID"`
	Usos        []UsoRecurso `gorm:"foreignKey:RecursoID"`
}

// RELACIONES

type AsignacionPersonal struct {
	IDAsignacion    uint      `json:"id_asignacion" gorm:"primaryKey;autoIncrement"`
	ProyectoID      uint      `json:"proyecto_id" gorm:"not null"`
	PersonalID      uint      `json:"personal_id" gorm:"not null"`
	HorasTrabajadas float64   `json:"horas_trabajadas" gorm:"not null"`
	FechaRegistro   time.Time `json:"fecha_registro" gorm:"not null"`

	Proyecto *Proyecto `gorm:"foreignKey:ProyectoID"`
	Personal *Personal `gorm:"foreignKey:PersonalID"`
	Pagos    []PagoPersonal `gorm:"foreignKey:AsignacionID"`
}

type UsoRecurso struct {
	IDUso          uint       `json:"id_uso" gorm:"primaryKey;autoIncrement"`
	ProyectoID     uint       `json:"proyecto_id" gorm:"not null"`
	RecursoID      uint       `json:"recurso_id" gorm:"not null"`
	FechaInicioUso time.Time  `json:"fecha_inicio_uso" gorm:"not null"`
	FechaFinUso    *time.Time `json:"fecha_fin_uso"`

	Proyecto *Proyecto      `gorm:"foreignKey:ProyectoID"`
	Recurso  *RecursoTecnico `gorm:"foreignKey:RecursoID"`
}

// CONTABLE

type Factura struct {
	IDFactura    uint      `json:"id_factura" gorm:"primaryKey;autoIncrement"`
	ProyectoID   *uint     `json:"proyecto_id"`
	ClienteID    *uint     `json:"cliente_id"`
	FechaEmision time.Time `json:"fecha_emision" gorm:"not null"`
	MontoTotal   float64   `json:"monto_total" gorm:"not null"`
	EstadoPagoID uint      `json:"estado_pago_id" gorm:"not null"`

	Proyecto   *Proyecto  `gorm:"foreignKey:ProyectoID"`
	Cliente    *Cliente   `gorm:"foreignKey:ClienteID"`
	EstadoPago *EstadoPago `gorm:"foreignKey:EstadoPagoID"`
}

type Gasto struct {
	IDGasto          uint      `json:"id_gasto" gorm:"primaryKey;autoIncrement"`
	ProyectoID       uint      `json:"proyecto_id" gorm:"not null"`
	DescripcionGasto string    `json:"descripcion_gasto" gorm:"size:255;not null"`
	CategoriaGastoID *uint     `json:"categoria_gasto_id"`
	MontoGasto       float64   `json:"monto_gasto" gorm:"not null"`
	FechaGasto       time.Time `json:"fecha_gasto" gorm:"not null"`

	Proyecto       *Proyecto       `gorm:"foreignKey:ProyectoID"`
	CategoriaGasto *CategoriaGasto `gorm:"foreignKey:CategoriaGastoID"`
}

type PagoPersonal struct {
	IDPago       uint      `json:"id_pago" gorm:"primaryKey;autoIncrement"`
	PersonalID   uint      `json:"personal_id" gorm:"not null"`
	ProyectoID   *uint     `json:"proyecto_id"`
	AsignacionID *uint     `json:"asignacion_id"`
	MontoPagado  float64   `json:"monto_pagado" gorm:"not null"`
	FechaPago    time.Time `json:"fecha_pago" gorm:"not null"`

	Personal   *Personal           `gorm:"foreignKey:PersonalID"`
	Proyecto   *Proyecto           `gorm:"foreignKey:ProyectoID"`
	Asignacion *AsignacionPersonal `gorm:"foreignKey:AsignacionID"`
}

// GESTOR (LOGIN)

type Gestor struct {
	ID        	uint      		`json:"id" gorm:"primaryKey;autoIncrement"`
	Nombre    	string    		`json:"nombre" gorm:"size:100;not null"`
	Usuario   	string    		`json:"usuario" gorm:"size:50;unique;not null"`
	Contrasena 	string   		`json:"-" gorm:"size:255;not null"`
	Email     	string    		`json:"email" gorm:"size:100;unique"`
	Telefono  	string    		`json:"telefono" gorm:"size:30"`
	CreatedAt 	time.Time 		`json:"created_at"`
	UpdatedAt 	time.Time 		`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleted_at" gorm:"index"`
}

type LoginRequest struct {
	Usuario string `json:"usuario"`
	Contrasena string `json:"contrasena"`
}

type GestorClaims struct {
	ID uint `json:"id"`
	Usuario string `json:"usuario"`
	jwt.RegisteredClaims
}
