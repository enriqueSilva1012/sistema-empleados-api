package models

import "time"

type Empleado struct {
	ID              uint `gorm:"primaryKey"`
	Nombre          string
	Apellido        string
	Puesto          string
	Salario         float64
	Departamento    string
	FechaIngreso    time.Time
	Estado          string
}