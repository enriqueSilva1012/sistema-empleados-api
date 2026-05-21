package database

import (
	"fmt"

	"proyecto-empleados/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarDB() {

	dsn := "root:@tcp(127.0.0.1:3306)/empleados_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error al conectar BD")
	}

	fmt.Println("Base de datos conectada")

	db.AutoMigrate(&models.Empleado{})

	DB = db
}