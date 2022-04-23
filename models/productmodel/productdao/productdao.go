package productdao

import (
	"fmt"
	"market_apis/internalservices/marketdb"
	"market_apis/models/productmodel"

	"gorm.io/gorm"
)

var (
	db *gorm.DB = marketdb.GetMarketDB().GetConnection()
)

// InsertProducts ..
func InsertProducts(record []productmodel.Product) error {

	result := db.Model(&productmodel.Product{}).Create(&record)
	if result.Error != nil {
		return fmt.Errorf(`Error when insert product into DB: %s`, result.Error.Error())
	}
	return nil
}

// FindProducts ..
func FindProducts(parameter interface{}) ([]productmodel.Product, error) {

	var products []productmodel.Product
	err := db.Model(&productmodel.Product{}).Where(parameter).Find(&products).Error
	if err != nil {
		return nil, fmt.Errorf("Error occure when get product from database: %s", err.Error())
	}
	return products, nil
}
