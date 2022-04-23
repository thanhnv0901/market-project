package testmodels

import (
	"gorm.io/gorm"
)

// Product ..
type Product struct {
	gorm.Model
	Name      *string  `json:"name" validate:"required,lte=100"`
	Quantity  *int32   `json:"quantity" gorm:"default:0" validate:"required,gte=1"`
	Unit      *string  `json:"unit" gorm:"default:peace"`
	Price     *float64 `json:"price" gorm:"default:0.0" validate:"required,gt=0"`
	PriceUnit *string  `json:"price_unit" gorm:"default:dollar"`
	UserID    *int32   `json:"user_id" gorm:"not null"`
	CompanyID *int32   `json:"company_id" gorm:"default:-1"`
}
