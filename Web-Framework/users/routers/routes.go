package routers

import (
	"github.com/labstack/echo"
	"github.com/marcelors/curso/usuarios/controllers"
)

// App e uma instancia de echo
var App *echo.Echo

func init() {
	App = echo.New()

	App.GET("/", controllers.Home)
	App.GET("/add", controllers.Add)
	App.GET("/edit/:id", controllers.Edit)

	api := App.Group("/users")
	api.POST("/", controllers.Save)
	api.DELETE("/:id", controllers.Delete)
	api.PUT("/:id", controllers.Update)
}
