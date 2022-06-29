package errors

import (
	"go-starter/response"
	"net/http"

	"gorm.io/gorm"
)

const (
	DATA_NOT_FOUND      = "data not found"
	FILE_NOT_FOUND      = "file not found"
	INVALID_FILE_FORMAT = "invalid file format"
	INVALID_PASSWORD    = "invalid password"
	PAYLOAD_TOO_LARGE   = "payload too large"
	PERMISSION_DENIED   = "permission denied"
)

func BadRequestException(w http.ResponseWriter, r *http.Request, err any) {
	switch err := err.(type) {
	case string:
		response.WriteJSON(w, r, response.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err,
		})
	default:
		response.WriteJSON(w, r, response.Response{
			StatusCode: http.StatusBadRequest,
			Error:      err.([]response.Error),
		})
	}
}

func NotFoundException(w http.ResponseWriter, r *http.Request, messages ...string) {
	message := DATA_NOT_FOUND
	if len(messages) > 0 {
		message = messages[0]
	}
	response.WriteJSON(w, r, response.Response{
		StatusCode: http.StatusNotFound,
		Message:    message,
	})
}

func InternalServerErrorException(w http.ResponseWriter, r *http.Request, message string) {
	response.WriteJSON(w, r, response.Response{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	})
}

func SqlError(w http.ResponseWriter, r *http.Request, err error) {
	switch err {
	case gorm.ErrRecordNotFound:
		NotFoundException(w, r, err.Error())
	default:
		InternalServerErrorException(w, r, err.Error())
	}
}
