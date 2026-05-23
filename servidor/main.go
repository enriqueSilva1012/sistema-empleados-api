package main

import (
	"proyecto-empleados/database"
	"proyecto-empleados/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	database.ConectarDB()

	r := gin.Default()

	routes.ConfigurarRutas(r)

	r.Run(":8080")
}
