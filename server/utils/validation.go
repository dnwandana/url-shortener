package utils

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validate(object interface{}) []string {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	validationErr := validate.Struct(object)

	if validationErr != nil {
		var errors []string
		for _, err := range validationErr.(validator.ValidationErrors) {
			translatedErr := fmt.Sprint(err.Translate(trans))
			errors = append(errors, translatedErr)
		}
		return errors
	}
	return nil
}
