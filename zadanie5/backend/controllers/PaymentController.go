package controllers

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var payments = map[string]*models.Payment{}

type PaymentController struct{}

func (productController *PaymentController) CreatePayment(context echo.Context) error {
	payment := new(models.Payment)
	if err := context.Bind(payment); err != nil {
		return err
	}

	payments[strconv.Itoa(int(payment.ID))] = payment
	return context.JSON(http.StatusCreated, payment)
}

func (productController *PaymentController) GetAllPayments(context echo.Context) error {
	allProducts := make([]*models.Payment, 0, len(payments))
	for _, payment := range payments {
		allProducts = append(allProducts, payment)
	}
	return context.JSON(http.StatusOK, allProducts)
}

func (productController *PaymentController) GetPayment(context echo.Context) error {
	payment := new(models.Payment)
	if err := context.Bind(payment); err != nil {
		return err
	}

	paymentID := context.Param("id")
	if payment, found := payments[paymentID]; found {
		return context.JSON(http.StatusOK, payment)
	}

	return context.JSON(http.StatusNotFound, "Payment not found")
}

func (productController *PaymentController) UpdatePayment(context echo.Context) error {
	payment := new(models.Payment)
	if err := context.Bind(payment); err != nil {
		return err
	}

	paymentID := context.Param("id")
	if _, found := payments[paymentID]; found {
		payments[paymentID] = payment
		return context.JSON(http.StatusOK, payment)
	}

	return context.JSON(http.StatusNotFound, "Payment not found")
}

func (productController *PaymentController) DeletePayment(context echo.Context) error {
	paymentID := context.Param("id")
	if _, found := payments[paymentID]; found {
		delete(payments, paymentID)
		return context.NoContent(http.StatusNoContent)
	}

	return context.JSON(http.StatusNotFound, "Payment not found")
}
