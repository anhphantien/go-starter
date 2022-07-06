package handlers

import (
	"go-starter/config"
	"go-starter/errors"
	"go-starter/response"
	"net/http"

	"golang.org/x/exp/slices"
)

type FileHandler struct{}

// @Tags    file
// @Summary Upload a file
// @Param   file                formData file false " "
// @Success 201                 object   response.Response{data=boolean}
// @Router  /api/v1/file/upload [POST]
func (h FileHandler) Upload(w http.ResponseWriter, r *http.Request) {
	_, fileHeader, err := r.FormFile("file")
	if err != nil {
		switch err {
		case http.ErrMissingBoundary:
			errors.BadRequestException(w, r, errors.FILE_NOT_FOUND)
		default:
			errors.BadRequestException(w, r, err.Error())
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

	// buffer := make([]byte, fileHeader.Size)
	// file.Read(buffer)
	// f, _ := os.Create("./" + fileHeader.Filename)
	// f.Write(buffer)

	response.WriteJSON(w, r, response.Response{
		Data: true,
	})
}
