package util

import (
	"encoding/json"
	"strings"
	"unicode"

	"github.com/go-playground/locales/pt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	pt_translations "github.com/go-playground/validator/v10/translations/pt"
)

var trans ut.Translator

type RequestError struct {
	Fields []Field `json:"fields"`
}

type Field struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r RequestError) Error() string {
	payload, _ := json.Marshal(r)
	return string(payload)
}

func NewValidator() *validator.Validate {
	pt := pt.New()
	uni := ut.New(pt, pt)

	trans, _ = uni.GetTranslator("pt")

	validate := validator.New()
	pt_translations.RegisterDefaultTranslations(validate, trans)

	return validate
}

func HandleValidatorFieldsError(err error) error {
	errs := err.(validator.ValidationErrors)

	requestError := RequestError{}

	for _, e := range errs {
		field := Field{
			Field:   getFormatedField(e.Namespace()),
			Message: e.Translate(trans),
		}

		requestError.Fields = append(requestError.Fields, field)
	}

	return requestError
}

func getFormatedField(field string) string {
	fieldSplited := strings.Split(field, ".")
	fieldConverted := ""

	for i, field := range fieldSplited {
		if i > 0 {
			fieldRune := []rune(field)
			fieldRune[0] = unicode.ToLower(fieldRune[0])
			field = string(fieldRune)

			fieldConverted += field

			if i < len(fieldSplited)-1 {
				fieldConverted += "."
			}
		}
	}

	return fieldConverted
}
