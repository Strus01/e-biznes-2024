package controllers

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var products = map[string]*models.Product{}

type ProductController struct{}

func (productController *ProductController) CreateProduct(context echo.Context) error {
	product := new(models.Product)
	if err := context.Bind(product); err != nil {
		return err
	}

	products[strconv.Itoa(int(product.ID))] = product
	return context.JSON(http.StatusCreated, product)
}

func (productController *ProductController) GetAllProducts(context echo.Context) error {
	allProducts := make([]*models.Product, 0, len(products))
	for _, product := range products {
		allProducts = append(allProducts, product)
	}
	return context.JSON(http.StatusOK, allProducts)
}

func (productController *ProductController) GetProduct(context echo.Context) error {
	product := new(models.Product)
	if err := context.Bind(product); err != nil {
		return err
	}

	productID := context.Param("id")
	if product, found := products[productID]; found {
		return context.JSON(http.StatusOK, product)
	}

	return context.JSON(http.StatusNotFound, "Product not found")
}

func (productController *ProductController) UpdateProduct(context echo.Context) error {
	product := new(models.Product)
	if err := context.Bind(product); err != nil {
		return err
	}

	productID := context.Param("id")
	if _, found := products[productID]; found {
		products[productID] = product
		return context.JSON(http.StatusOK, product)
	}

	return context.JSON(http.StatusNotFound, "Product not found")
}

func (productController *ProductController) DeleteProduct(context echo.Context) error {
	productID := context.Param("id")
	if _, found := products[productID]; found {
		delete(products, productID)
		return context.NoContent(http.StatusNoContent)
	}

	return context.JSON(http.StatusNotFound, "Product not found")
}
