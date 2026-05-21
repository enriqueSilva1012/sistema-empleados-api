package main

import (
	"proyecto-empleados/database"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConectarDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"mensaje":"Servidor funcionando",
		})

	})

	r.Run(":8080")

}