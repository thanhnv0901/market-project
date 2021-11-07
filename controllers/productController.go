package controllers

import (
	"market_apis/handlers"
	"market_apis/internals/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ProductionController ..
type ProductionController struct {
}

// NewProductionController ..
func NewProductionController() *ProductionController {
	return &ProductionController{}
}

func responce(c echo.Context, statusCode int, message string, isSuccess bool) error {
	return c.JSON(statusCode, map[string]interface{}{
		"Success": isSuccess,
		"Message": message,
	})
}

// UploadProduct ..
func (p *ProductionController) UploadProduct(c echo.Context) error {
	defer utils.ErrorTrackingDefer()

	var (
		productHandler handlers.ProductHandler = *handlers.NewProductHandler()
		err            error
	)

	err = productHandler.InsertProduct(c)
	if err != nil {
		c.Logger().Errorf("Cannot insert product: %s\n", err.Error())
		return responce(c, http.StatusBadRequest, err.Error(), false)
	}

	return responce(c, http.StatusOK, "OK", true)
}
