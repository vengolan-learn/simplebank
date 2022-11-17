package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/vengolan/simplebank/db/util"
)

var validCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		//check currency imported
		return util.IsSupportedCurrency(currency)
	}
	return false

}
