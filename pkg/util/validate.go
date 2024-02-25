package util

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/pkg/errors"
	"reflect"
	"strings"
)

type Validator struct {
	*validator.Validate
	Translator *ut.Translator
}

func NewValidator() (*Validator, error) {

	translator := en.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("en")
	if !found {
		return nil, errors.New("translator not found")
	}

	validate := validator.New()
	if validate == nil {
		return nil, errors.New("unable to create validator instance")
	}

	// register default translations
	if err := enTranslations.RegisterDefaultTranslations(validate, trans); err != nil {
		return nil, errors.WithMessage(err, "unable to register en translations")
	}

	// custom message for uuid_rfc4122
	_ = validate.RegisterTranslation("uuid_rfc4122", trans, func(ut ut.Translator) error {
		return ut.Add("uuid_rfc4122", "{0} must be a uuid", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("uuid_rfc4122", fe.Field())
		return t
	})

	// use the exact name instead of the capitalized
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return &Validator{
		Validate:   validate,
		Translator: &trans,
	}, nil
}
