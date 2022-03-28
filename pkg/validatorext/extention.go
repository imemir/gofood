package validatorext

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
)

const (
	mobileRegexString = `^(\+98|0)?9\d{9}$`
)

var (
	validate    = validator.New()
	mobileRegex = regexp.MustCompile(mobileRegexString)
)

func init() {
	_ = validate.RegisterValidation("mobile", ValidateMobile)
}

type EnumValidation interface {
	InRange(v interface{}) bool
}

func ValidateMobile(fl validator.FieldLevel) bool {
	if fl.Field().Kind() != reflect.String {
		return false
	}
	v := fl.Field().String()
	return mobileRegex.MatchString(v)
}

func Struct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		return err
	}
	return nil
}
