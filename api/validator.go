package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/stanely158831384/guluguluStorage/util"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool{
	if currency, ok := fieldLevel.Field().Interface().(string); ok{
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}

var validCategory validator.Func = func(fieldLevel validator.FieldLevel) bool{
	if category, ok := fieldLevel.Field().Interface().(string); ok{
		// check category is supported
		return util.CategoryDetector(category)
	}
	return false
}