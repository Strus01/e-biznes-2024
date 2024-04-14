package routes

import (
	"backend/controllers"

	"github.com/labstack/echo/v4"
)

func PaymentRoute(e *echo.Echo) {
	paymentController := controllers.PaymentController{}

	e.POST("/payments", paymentController.CreatePayment)
	e.GET("/payments", paymentController.GetAllPayments)
	e.GET("/payments/:id", paymentController.GetPayment)
	e.PUT("/payments/:id", paymentController.UpdatePayment)
	e.DELETE("/payments/:id", paymentController.DeletePayment)
}
