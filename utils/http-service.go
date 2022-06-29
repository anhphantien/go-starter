package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HttpGet() (map[string]any, error) {
	// buffer, _ := json.Marshal(map[string]any{
	// 	"username": "superadmin",
	// 	"password": "123456",
	// })
	// req, _ := http.NewRequest(http.MethodPost, "<url>", bytes.NewBuffer(buffer))
	client := http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "<url>", nil)
	req.Header = http.Header{
		"Content-Type": {"application/json"},
		// "Authorization": {"Bearer <token>"},
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	buffer, _ := io.ReadAll(res.Body)
	fmt.Println(string(buffer))
	data := map[string]any{}
	err = json.Unmarshal([]byte(string(buffer)), &data)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	return data, nil
}
