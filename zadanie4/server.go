package main

import (
	"zadanie4/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	routes.ProductRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
