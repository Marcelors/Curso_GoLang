package main

import (
	"github.com/MarcusMann/pongor-echo"
	"github.com/labstack/echo/middleware"
	r "github.com/marcelors/curso/usuarios/routers"
)

func main() {
	e := r.App
	e.Use(middleware.Logger())

	p := pongor.GetRenderer()
	p.Directory = "views"
	e.Renderer = p

	e.Logger.Fatal(e.Start(":8000"))
}
