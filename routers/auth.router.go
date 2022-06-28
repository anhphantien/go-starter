package routers

import (
	"go-starter/handlers"
	"net/http"
)

func AuthRouter() {
	http.HandleFunc("/auth/login", handlers.AuthLogin)
	http.HandleFunc("/users/profile", handlers.UserProfile)
}
