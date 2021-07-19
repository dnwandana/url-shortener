package util

import (
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	universalTranslator *ut.UniversalTranslator
	validate            *validator.Validate
)

// Validate struct level.
func Validate(object interface{}) interface{} {
	// declare universal-translator
	enTranslate := en.New()
	universalTranslator = ut.New(enTranslate, enTranslate)
	trans, _ := universalTranslator.GetTranslator("en_translator")

	// new validator instance
	validate = validator.New()
	// register translator
	errTranslation := en_translations.RegisterDefaultTranslations(validate, trans)
	if errTranslation != nil {
		return errTranslation.Error()
	}

	// validate Struct level
	validationErr := validate.Struct(object)
	// if there is an error
	if validationErr != nil {
		for _, err := range validationErr.(validator.ValidationErrors) {
			translatedErr := fmt.Sprintf(err.Translate(trans))
			return translatedErr
		}
	}
	return nil
}
