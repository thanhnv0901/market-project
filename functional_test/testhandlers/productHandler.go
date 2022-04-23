package testhandlers

import (
	"fmt"
	"market_apis/functional_test/testmodels"
)

// ProductHandler ..
type ProductHandler struct {
}

// TruncateTable ..
func (p *ProductHandler) TruncateTable() error {
	marketDBConnection.Unscoped().Delete(&testmodels.Product{}, "1=1")
	return nil
}

// GetAllData ..
func (p *ProductHandler) GetAllData() ([]map[string]interface{}, error) {
	rows := make([]map[string]interface{}, 0)
	result := marketDBConnection.Table(ProductTable).Find(&rows)
	err := result.Error
	if err != nil {
		return rows, fmt.Errorf("Error occure when get data in table Products: %s", err.Error())
	}
	return rows, nil
}

// InsertData ..
func (p *ProductHandler) InsertData(rows []map[string]interface{}) error {

	result := marketDBConnection.Table(ProductTable).Create(&rows)
	err := result.Error
	if err != nil {
		return fmt.Errorf("Error occure when get data in table Products: %s", err.Error())
	}
	return nil
}
