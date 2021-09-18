package controllers

import (
	"market_apis/internals/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// ProductionController ..
type ProductionController struct {
}

// NewProductionController ..
func NewProductionController() *ProductionController {
	return &ProductionController{}
}

// UploadProduct ..
func (p *ProductionController) UploadProduct(c echo.Context) error {
	defer utils.ErrorTrackingDeder()

	// panic("error roi do")

	c.Logger().Errorj(log.JSON{
		"message": " aok",
	})

	return c.JSON(http.StatusOK, map[string]string{
		"Hello": "OKE",
	})
}
