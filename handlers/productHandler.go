package handlers

import (
	"market_apis/models"

	"github.com/labstack/echo/v4"
)

// ProductHandler ..
type ProductHandler struct {
}

// NewProductHandler ..
func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// InsertProduct ..
func (p *ProductHandler) InsertProduct(c echo.Context) (err error) {

	var product models.Product
	err = c.Bind(&product)
	if err != nil {
		return err
	}
	err = product.Validate()
	if err != nil {
		return err
	}

	err = product.InsertProduct()
	if err != nil {
		return err
	}
	return nil
}
