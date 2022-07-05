package handlers

import (
	"fmt"
	"go-starter/errors"
	"net/http"
)

type FileHandler struct{}

// @Tags    file
// @Summary Upload a file
// @Param   file                formData file false " "
// @Success 201                 object   response.Response{}
// @Router  /api/v1/file/upload [POST]
func (h FileHandler) Upload(w http.ResponseWriter, r *http.Request) {
	_, header, err := r.FormFile("file")
	if err != nil { // http.ErrMissingBoundary
		errors.BadRequestException(w, r, errors.FILE_NOT_FOUND)
	}
	fmt.Println(header)
}
