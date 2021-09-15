package util

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

// Validate struct, returns slice of string error.
func Validate(object interface{}) []string {
	// declare universal-translator
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	// new validator instance
	validate = validator.New()
	// register translator
	en_translations.RegisterDefaultTranslations(validate, trans)

	// validate Struct level
	validationErr := validate.Struct(object)
	// if there is an error
	// returns slice of string error
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
