package utils

import (
	"encoding/json"
	"go-starter/response"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRequestBody(r *http.Request, payload any) []response.Error {
	_payload, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(_payload, payload)

	if err := validator.New().Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make([]response.Error, len(validationErrors))

		for i, fieldError := range validationErrors {
			errors[i] = response.Error{
				Field: strings.ToLower(fieldError.Field()),
				Message: func(fieldError validator.FieldError) string {
					switch fieldError.Tag() {
					case "required":
						return "This field is required"
					case "max":
						return "Max length: " + fieldError.Param()
					default:
						return fieldError.Error()
					}
				}(fieldError),
			}
		}
		return errors
	}
	return nil
}
