package validators

import (
	"backend-master-class/util"

	"github.com/go-playground/validator/v10"
)

var ValidCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(int32); ok {
		return util.IsSupportedCurrency(currency)
	}

	return false
}
