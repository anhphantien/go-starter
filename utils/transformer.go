package utils

import (
	"encoding/json"
	"strconv"
)

func ConvertToID(v any) uint64 {
	switch _v := v.(type) {
	case *uint64:
		return *_v
	case float64:
		return uint64(v.(float64))
	case string:
		n, _ := strconv.ParseUint(v.(string), 10, 64)
		return n
	default:
		return v.(uint64)
	}
}

func ConvertToMap(v any) map[string]any {
	m := map[string]any{}
	_v, _ := json.Marshal(v)
	json.Unmarshal(_v, &m)
	return m
}
