package main

import (
	_ "go-starter/docs"
	"go-starter/routers"
	"net/http"
)

// @title Go starter
// @version 1.0
// @description Go starter's API documentation
// @BasePath /api
func main() {
	routers.New()

	http.ListenAndServe(":8000", nil)
}
