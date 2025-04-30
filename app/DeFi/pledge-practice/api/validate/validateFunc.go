package validate

import (
	"github.com/go-playground/validator/v10"
	"p-cy/util"
	"regexp"
)

func IsPassword(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		if isOk, _ := regexp.MatchString(`^[a-zA-Z0-9!@#￥%^&*]{6,20}$`, fl.Field().Interface().(string)); isOk {
			return isOk
		}
	}
	return false
}

func IsEmail(fl validator.FieldLevel) bool {
	if fl.Field().Interface().(string) != "" {
		if isOk, _ := regexp.MatchString(util.EMAIL_REG, fl.Field().Interface().(string)); isOk {
			return isOk
		}
	}
	return false
}
