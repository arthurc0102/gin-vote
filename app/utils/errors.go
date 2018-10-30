package utils

import (
	"encoding/json"
	"fmt"
	"reflect"

	validator "gopkg.in/go-playground/validator.v8"
)

// FieldValidatorError change validator.FieldError to string
func FieldValidatorError(e *validator.FieldError) string {
	switch e.Tag {
	case "required":
		return fmt.Sprint("This field is required")
	case "notspace":
		return fmt.Sprint("This field is required and can't be space")
	case "max":
		return fmt.Sprintf("This field cannot be bigger or longer than %s", e.Param)
	case "min":
		return fmt.Sprintf("This field must be bigger or longer than %s", e.Param)
	case "email":
		return fmt.Sprint("Invalid email format")
	case "len":
		return fmt.Sprintf("This field must be %s characters long", e.Param)
	}

	return fmt.Sprint("This field is not valid")
}

// HandleError handle binding error
func HandleError(object interface{}, err error) map[string]string {
	errors := make(map[string]string)
	val := reflect.ValueOf(object)

	switch reflect.TypeOf(err) {
	case reflect.TypeOf(validator.ValidationErrors{}):
		for _, e := range err.(validator.ValidationErrors) {
			field, _ := val.Type().FieldByName(e.Field)
			jsonKey := field.Tag.Get("json")

			if jsonKey == "" {
				jsonKey = e.Field
			}

			errors[jsonKey] = FieldValidatorError(e)
		}
	case reflect.TypeOf(&json.SyntaxError{}):
		errors["nonField"] = "Parse json error"
	case reflect.TypeOf(&json.UnmarshalTypeError{}):
		err := err.(*json.UnmarshalTypeError)
		errors[err.Field] = fmt.Sprintf("This field required a %s type value, not %s", err.Type, err.Value)
	default:
		errors["nonField"] = err.Error()
	}

	return errors
}
