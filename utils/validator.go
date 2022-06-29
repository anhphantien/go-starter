package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Error struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateRequestBody(r *http.Request, body any) any {
	_body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal([]byte(_body), body)

	if err := validator.New().Struct(body); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errors := make([]Error, len(validationErrors))

		for i, fieldError := range validationErrors {
			errors[i] = Error{
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

// func FilterRequestBody(c *fiber.Ctx, body any) map[string]any {
// 	dto := map[string]any{}
// 	_dto, _ := json.Marshal(body)
// 	json.Unmarshal(_dto, &dto)

// 	rawBody := map[string]any{}
// 	json.Unmarshal(c.Body(), &rawBody)

// 	return lo.PickByKeys(rawBody, maps.Keys(dto))
// }
