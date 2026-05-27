package models

import (
	"time"

	"gorm.io/gorm"
)

type Empleado struct {
	gorm.Model

	Nombre string `json:"nombre"`

	Puesto string `json:"puesto"`

	Departamento string `json:"departamento"`

	FechaIngreso time.Time `json:"fechaIngreso"`
}
