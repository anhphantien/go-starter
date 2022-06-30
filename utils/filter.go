package utils

import (
	"encoding/json"
	"net/http"

	"github.com/samber/lo"
	"golang.org/x/exp/maps"
)

func FilterRequestBody(r *http.Request, payload any) map[string]any {
	dto := map[string]any{}
	_dto, _ := json.Marshal(payload)
	json.Unmarshal(_dto, &dto)

	rawBody := map[string]any{}

	return lo.PickByKeys(rawBody, maps.Keys(dto))
}
