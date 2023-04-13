package dto

import (
	"outgrow/enum"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()

	err := validate.RegisterValidation("er_amount", ValidateEventRuleAmount)
	if err != nil {
		panic(err)
	}
}

func Validate(s interface{}) (err error) {
	if reflect.TypeOf(s).Kind() == reflect.Slice {
		err = validate.Var(s, "dive")
	} else {
		err = validate.Struct(s)
	}

	return
}

func ValidateEventRuleAmount(fl validator.FieldLevel) bool {
	v := fl.Field().Float()
	ruleType := fl.Parent().FieldByName("RuleType").Interface().(enum.EventRuleType)
	if ruleType == enum.EventRulesPercentage && v > 100 {
		return false
	}

	return true
}
