package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) map[string][]string {
	result := make(map[string][]string)
	fmt.Println(err.Error())

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		e := strings.Split(err.Error(), ".")[1]
		errors := strings.Split(e, " of type ")
		field := ToSnakeCase(errors[0])
		message := fmt.Sprintf("%s must be %s", field, errors[1])
		if val, ok := result[field]; ok {
			val = append(val, message)
			result[field] = val
		} else {
			result[field] = []string{message}
		}
		return result
	}

	if !errors.Is(err, io.EOF) {
		for _, e := range err.(validator.ValidationErrors) {
			field := ToSnakeCase(e.Field())
			var message string
			if e.Param() != "" {
				message = ToSnakeCase(e.Field()) + " " + e.Tag() + " " + e.Param()
			} else {
				message = ToSnakeCase(e.Field()) + " " + e.Tag()
			}
			if val, ok := result[field]; ok {
				val = append(val, message)
				result[field] = val
			} else {
				result[field] = []string{message}
			}
		}
	}
	return result
}
