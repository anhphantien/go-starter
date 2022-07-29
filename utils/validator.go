package utils

import (
	"encoding/json"
	"go-starter/errors"
	"go-starter/response"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateRequestBody(w http.ResponseWriter, r *http.Request, payload any) []response.Error {
	_payload, _ := io.ReadAll(r.Body)
	json.Unmarshal(_payload, payload)

	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		err := make([]response.Error, len(validationErrors))

		for i, fieldError := range validationErrors {
			err[i] = response.Error{
				Field: fieldError.Field(),
				Message: func(fe validator.FieldError) string {
					switch fe.Tag() {
					case "required":
						return "This field is required"
					case "max":
						return "Max length: " + fe.Param()
					default:
						return fe.Error()
					}
				}(fieldError),
			}
		}

		errors.BadRequestException(w, r, err)
		return err
	}
	return nil
}
