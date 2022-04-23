package handlers

import (
	"fmt"
	"market_apis/models/productmodel"
	"market_apis/models/productmodel/productdao"

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

	var product productmodel.Product
	err = c.Bind(&product)
	if err != nil {
		return fmt.Errorf("Error occure when bind request body to object: %s", err.Error())
	}
	err = product.Validate()
	if err != nil {
		return err
	}
	tmp := []productmodel.Product{product}
	err = productdao.InsertProducts(tmp)
	if err != nil {
		return fmt.Errorf("Error occure when insert into DB: %s", err.Error())
	}

	return nil
}

// GetProductsByAtribute ..
func (p *ProductHandler) GetProductsByAtribute(parameter interface{}) (products []productmodel.Product, err error) {

	products, err = productdao.FindProducts(parameter)
	if err != nil {
		return nil, fmt.Errorf("Error occure when product by attribute: %s", err.Error())
	}
	return products, nil
}
