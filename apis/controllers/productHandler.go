package controllers

import "github.com/labstack/echo/v4"

// ProductionController ..
type ProductionController struct {
}

// NewProductionController ..
func NewProductionController() *ProductionController {
	return &ProductionController{}
}

// SellProduct ..
func (p *ProductionController) SellProduct(c echo.Context) {

}
