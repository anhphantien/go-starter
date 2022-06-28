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
	r := routers.New()
	log.Fatal(http.ListenAndServe(":"+env.PORT, r))
}
