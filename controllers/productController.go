package controllers

import (
	"market_apis/handlers"
	"market_apis/internals/utils"
	"market_apis/models"
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

func responce(c echo.Context, statusCode int, message string, isSuccess bool, data interface{}) error {
	return c.JSON(statusCode, map[string]interface{}{
		"success": isSuccess,
		"message": message,
		"data":    data,
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
		return responce(c, http.StatusBadRequest, err.Error(), false, nil)
	}

	return responce(c, http.StatusOK, "OK", true, nil)
}

type productQueryParameter struct {
	ID   int    `query:"id"`
	Name string `query:"name"`
}

// GetProduct ..
func (p *ProductionController) GetProduct(c echo.Context) error {

	var (
		productHandler   handlers.ProductHandler = *handlers.NewProductHandler()
		err              error
		productParameter productQueryParameter
	)
	c.Bind(&productParameter)

	var respData []models.Product
	respData, err = productHandler.GetProductsByAtribute(productParameter)
	if err != nil {
		c.Logger().Errorf("Cannot get product: %s\n", err.Error())
		return responce(c, http.StatusBadRequest, err.Error(), false, nil)
	}
	return responce(c, http.StatusOK, "OK", true, respData) // notice
}
