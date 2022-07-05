package handlers

import (
	"fmt"
	"go-starter/config"
	"go-starter/errors"
	"net/http"

	"golang.org/x/exp/slices"
)

type FileHandler struct{}

// @Tags    file
// @Summary Upload a file
// @Param   file                formData file false " "
// @Success 201                 object   response.Response{}
// @Router  /api/v1/file/upload [POST]
func (h FileHandler) Upload(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		switch err {
		case http.ErrMissingBoundary:
			errors.BadRequestException(w, r, errors.FILE_NOT_FOUND)
		default:
			errors.BadRequestException(w, r, err)
		}
		return
	}

	if !slices.Contains(
		[]string{
			config.File.ContentType.JPEG,
			config.File.ContentType.PNG,
		}, fileHeader.Header["Content-Type"][0]) {
		errors.BadRequestException(w, r, errors.INVALID_FILE_FORMAT)
		return
	}

	if fileHeader.Size > config.File.MaxSize {
		errors.PayloadTooLargeException(w, r)
		return
	}

	fmt.Println(file, fileHeader.Filename)
}
