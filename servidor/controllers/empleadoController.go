package controllers

import (
	"proyecto-empleados/database"
	"proyecto-empleados/models"

	"github.com/gin-gonic/gin"
)

func ObtenerEmpleados(c *gin.Context) {

	var empleados []models.Empleado

	database.DB.Find(&empleados)

	c.JSON(200, empleados)

}

func ObtenerEmpleado(c *gin.Context) {

	id := c.Param("id")

	var empleado models.Empleado

	database.DB.First(&empleado, id)

	c.JSON(200, empleado)

}

func CrearEmpleado(c *gin.Context) {

	var empleado models.Empleado

	if err := c.ShouldBindJSON(&empleado); err != nil {

		c.JSON(400, gin.H{
			"error": err.Error(),
		})

		return

	}

	database.DB.Create(&empleado)

	c.JSON(201, empleado)

}

func ActualizarEmpleado(c *gin.Context) {

	id := c.Param("id")

	var empleado models.Empleado

	database.DB.First(&empleado, id)

	var datos models.Empleado

	c.ShouldBindJSON(&datos)

	database.DB.Model(&empleado).Updates(datos)

	c.JSON(200, empleado)

}

func EliminarEmpleado(c *gin.Context) {

	id := c.Param("id")

	var empleado models.Empleado

	database.DB.First(&empleado, id)

	database.DB.Delete(&empleado)

	c.JSON(
		200,
		gin.H{
			"mensaje": "Empleado eliminado",
		},
	)

}
