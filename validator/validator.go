package validator

import (
	"gopkg.in/go-playground/validator.v8"
)

var Validate *validator.Validate

func SetValidator(tagName string) {
	config := &validator.Config{TagName: tagName}
	Validate = validator.New(config)
}
