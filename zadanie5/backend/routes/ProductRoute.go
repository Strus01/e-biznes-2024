package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

func ProductRoute(e *echo.Echo) {
	productController := controllers.ProductController{}

	e.POST("/products", productController.CreateProduct)
	e.GET("/products", productController.GetAllProducts)
	e.GET("/products/:id", productController.GetProduct)
	e.PUT("/products/:id", productController.UpdateProduct)
	e.DELETE("/products/:id", productController.DeleteProduct)
}
