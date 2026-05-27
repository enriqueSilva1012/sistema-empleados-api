package main

import (
	"proyecto-empleados/database"
	"proyecto-empleados/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	database.ConectarDB()

	r := gin.Default()

	r.Use(cors.Default())

	routes.ConfigurarRutas(r)

	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"mensaje": "Servidor REST API funcionando 🚀",
		})

	})

	// Permite conexiones externas
	r.Run("0.0.0.0:8080")
}
