package main

import (
	_ "go-starter/docs"
	"go-starter/env"
	"go-starter/routers"
	"log"
	"net/http"
)

// @title Go starter
// @version 1.0
// @description Go starter's API documentation
func main() {
	prefix := "/api/v1"
	routers.New(prefix)

	log.Fatal(http.ListenAndServe(":"+env.PORT, nil))
}
