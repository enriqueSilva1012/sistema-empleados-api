package models

import "time"

type Empleado struct {
	ID uint `gorm:"primaryKey" json:"id"`

	Nombre string `json:"nombre"`

	Apellido string `json:"apellido"`

	Puesto string `json:"puesto"`

	Salario float64 `json:"salario"`

	Departamento string `json:"departamento"`

	FechaIngreso time.Time `json:"fechaIngreso"`

	Estado string `json:"estado"`
}
