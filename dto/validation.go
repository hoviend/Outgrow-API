package dto

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(s interface{}) (err error) {
	if reflect.TypeOf(s).Kind() == reflect.Slice {
		err = validate.Var(s, "dive")
	} else {
		err = validate.Struct(s)
	}

	return
}
