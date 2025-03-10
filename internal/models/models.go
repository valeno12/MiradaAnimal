// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package models

import (
	"database/sql"
	"time"
)

type Animal struct {
	ID        uint64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Nombre    string
	Especie   string
	Foto      sql.NullString
	Sexo      string
	EdadMeses int32
	DeletedAt sql.NullTime
}

type AnimalAtributo struct {
	ID         uint64
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
	AnimalID   int32
	AtributoID int32
	DeletedAt  sql.NullTime
}

type AnimalEstado struct {
	ID        uint64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	PersonaID sql.NullInt32
	AnimalID  int32
	EstadoID  int32
	Final     sql.NullTime
	DeletedAt sql.NullTime
}

type AnimalTurno struct {
	ID            uint64
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
	Descripcion   string
	Fecha         sql.NullTime
	AnimalID      int32
	VeterinarioID int32
	DeletedAt     sql.NullTime
}

type AnimalVacuna struct {
	ID        uint64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	AnimalID  int32
	VacunaID  int32
	Fecha     time.Time
	DeletedAt sql.NullTime
}

type Atributo struct {
	ID          uint64
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	Descripcion string
	DeletedAt   sql.NullTime
}

type CitaVeterinario struct {
	ID            uint64
	CreatedAt     sql.NullTime
	UpdatedAt     sql.NullTime
	Descripcion   string
	Fecha         time.Time
	Motivo        string
	Peso          sql.NullString
	Costo         sql.NullString
	AnimalID      int32
	VeterinarioID int32
	DeletedAt     sql.NullTime
}

type Clinica struct {
	ID        uint64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Nombre    string
	Direccion string
	Telefono  string
	DeletedAt sql.NullTime
}

type Estado struct {
	ID        uint64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Nombre    string
	DeletedAt sql.NullTime
}

type Formulario struct {
	ID        uint64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	AnimalID  int32
	PersonaID int32
	DeletedAt sql.NullTime
}

type Permiso struct {
	ID          uint64
	Nombre      string
	Descripcion string
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	DeletedAt   sql.NullTime
}

type Persona struct {
	ID        uint64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Dni       string
	Nombre    string
	Apellido  string
	Direccion string
	Telefono  string
	Email     sql.NullString
	Estado    string
	DeletedAt sql.NullTime
}

type PersonaDetalle struct {
	ID         uint64
	CreatedAt  sql.NullTime
	UpdatedAt  sql.NullTime
	PersonaID  int32
	PreguntaID int32
	Detalle    sql.NullString
	DeletedAt  sql.NullTime
}

type Pregunta struct {
	ID              uint64
	CreatedAt       sql.NullTime
	UpdatedAt       sql.NullTime
	Descripcion     string
	RequiereDetalle bool
	DeletedAt       sql.NullTime
}

type Usuario struct {
	ID        uint64
	Nombre    string
	Email     string
	Password  string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type UsuarioPermiso struct {
	ID        uint64
	UsuarioID int32
	PermisoID int32
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}

type Vacuna struct {
	ID          uint64
	CreatedAt   sql.NullTime
	UpdatedAt   sql.NullTime
	Nombre      string
	Descripcion string
	DeletedAt   sql.NullTime
}

type Veterinario struct {
	ID        uint64
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	Dni       string
	Nombre    string
	Apellido  string
	Telefono  string
	ClinicaID int32
	DeletedAt sql.NullTime
}
