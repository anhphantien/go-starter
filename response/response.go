package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Property string `json:"property"`
	Message  string `json:"message"`
}

type Response struct {
	StatusCode int     `json:"statusCode"`
	Data       any     `json:"data,omitempty"`
	Message    string  `json:"message,omitempty"`
	Error      []Error `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, r *http.Request, payload Response) {
	if payload.StatusCode == 0 {
		if r.Method == http.MethodPost {
			payload.StatusCode = http.StatusCreated
		} else {
			payload.StatusCode = http.StatusOK
		}
	}
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(payload.StatusCode)
	w.Write(res)
}
