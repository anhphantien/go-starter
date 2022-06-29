package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type errorResp struct {
	Error string `json:"error"`
}

type meResp struct {
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

type authResp struct {
	Token string `json:"token"`
}

type authParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct{}

// WriteJSON provides function to format output response in JSON
func WriteJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error Parsing JSON")
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

// BasicAuthMW is middleware function to check whether user is authenticated or not
// actually you could write better code for this function
func BasicAuthMW(w http.ResponseWriter, r *http.Request) map[string]string {
	errorAuth := errorResp{
		Error: "Unauthorized access",
	}

	header := r.Header.Get("Authorization")
	if header == "" {
		WriteJSON(w, 401, errorAuth)
		return map[string]string{}
	}

	apiKey := strings.Split(header, " ")

	if len(apiKey) != 2 {
		WriteJSON(w, 401, errorAuth)
		return map[string]string{}
	}

	if apiKey[0] != "Basic" {
		WriteJSON(w, 401, errorAuth)
		return map[string]string{}
	}

	users := map[string]map[string]string{
		"28b662d883b6d76fd96e4ddc5e9ba780": {
			"username": "linggar",
			"fullname": "Linggar Primahastoko",
		},
	}

	if _, ok := users[apiKey[1]]; !ok {
		WriteJSON(w, 401, errorAuth)
		return map[string]string{}
	}

	return users[apiKey[1]]
}

func DecodePost(r *http.Request, structure interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(structure)
	if err != nil {
		log.Println("Error parsing post data")
	}
}

// @Summary Auth Login
// @Description Auth Login
// @Tags auth
// @ID auth-login
// @Accept  json
// @Produce  json
// @Param AuthLogin body authParam true "Auth Login Input"
// @Success 200 {object} authResp
// @Router /api/v1/auth/login [post]
func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var param authParam
	DecodePost(r, &param)

	if param.Username == "linggar" && param.Password == "linggar" {
		respAuth := authResp{
			Token: "28b662d883b6d76fd96e4ddc5e9ba780",
		}
		WriteJSON(w, 200, respAuth)
	} else {
		failResp := errorResp{
			Error: "Wrong username/password",
		}
		WriteJSON(w, 401, failResp)
	}
}

// UserProfile godoc
// @Summary User Profile
// @Description User Profile
// @Tags users
// @ID user-profile
// @Accept  json
// @Produce  json
// @Success 200 {object} meResp
// @Router /api/v1/users/profile [get]
// @Security Bearer
func UserProfile(w http.ResponseWriter, r *http.Request) {
	info := BasicAuthMW(w, r)

	if len(info) == 0 {
		return
	}

	respMe := meResp{
		Username: info["username"],
		Fullname: info["fullname"],
	}

	WriteJSON(w, 200, respMe)
}
