package routes

import (
	"proyecto-empleados/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigurarRutas(r *gin.Engine) {

	r.GET(
		"/empleados",
		controllers.ObtenerEmpleados,
	)

	r.GET(
		"/empleados/:id",
		controllers.ObtenerEmpleado,
	)

	r.POST(
		"/empleados",
		controllers.CrearEmpleado,
	)

	r.PUT(
		"/empleados/:id",
		controllers.ActualizarEmpleado,
	)

	r.DELETE(
		"/empleados/:id",
		controllers.EliminarEmpleado,
	)

	r.GET(
		"/departamento/:departamento",
		controllers.BuscarDepartamento,
	)
}
