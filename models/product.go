package models

import (
	"fmt"
	"market_apis/internalservices/marketdb"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB        = marketdb.GetMarketDB().GetConnection()
	productValidator ValidateProduct = *NewValidateProduct("us")
)

func init() {

}

// ValidateProduct ..
type ValidateProduct struct {
	Validator *validator.Validate
	trans     ut.Translator
	language  string
}

// Check ..
func (v *ValidateProduct) Check(product interface{}) error {

	var errMessages []string

	err := v.Validator.Struct(product)
	if err == nil {
		return nil
	}
	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		errMessages = append(errMessages, e.Translate(v.trans))
	}

	msg := strings.Join(errMessages, ", ")
	return fmt.Errorf(msg)
}

// NewValidateProduct ..
func NewValidateProduct(language string) *ValidateProduct {

	validate := validator.New()
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator(language)
	en_translations.RegisterDefaultTranslations(validate, trans)

	validateProduct := ValidateProduct{
		Validator: validate,
		language:  language,
		trans:     trans,
	}

	return &validateProduct
}

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

// Validate ..
func (p *Product) Validate() error {
	return productValidator.Check(p)
}

// InsertProduct ..
func (p *Product) InsertProduct() error {
	result := db.Create(&p)
	return result.Error
}
