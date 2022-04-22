package handlers

import (
	"fmt"
	"market_apis/internalservices/marketdb"
	"market_apis/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = marketdb.GetMarketDB().GetConnection()
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
		return fmt.Errorf("Error occure when bind request body to object: %s", err.Error())
	}
	err = product.Validate()
	if err != nil {
		return err
	}

	result := db.Create(&product)
	err = result.Error
	if err != nil {
		return fmt.Errorf("Error occure when insert into DB: %s", err.Error())
	}

	return nil
}

// GetProductsByAtribute ..
func (p *ProductHandler) GetProductsByAtribute(parameter interface{}) ([]models.Product, error) {

	var products []models.Product
	err := db.Find(&products, parameter).Error
	if err != nil {
		return nil, fmt.Errorf("Error occure when file product: %s", err.Error())
	}
	return products, nil
}
