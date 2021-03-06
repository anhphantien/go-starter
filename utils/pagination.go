package utils

import (
	"encoding/json"
	"go-starter/dto"
	"net/http"
	"strconv"

	"golang.org/x/exp/slices"
)

func Pagination(r *http.Request) dto.Pagination {
	query := r.URL.Query()

	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit < 1 || limit > 100 {
		limit = 10
	}

	page, _ := strconv.Atoi(query.Get("page"))
	if page < 1 {
		page = 1
	}

	keyword := query.Get("keyword")

	filter := map[string]any{}
	json.Unmarshal([]byte(query.Get("filter")), &filter)

	var sort struct {
		By        string
		Direction string
	}
	json.Unmarshal([]byte(query.Get("sort")), &sort)
	if len(sort.By) == 0 {
		sort.By = "id"
	}
	if !slices.Contains(
		[]string{
			"ASC",
			"DESC",
		}, sort.Direction) {
		sort.Direction = "DESC"
	}

	return dto.Pagination{
		Limit:   limit,
		Offset:  limit * (page - 1),
		Keyword: keyword,
		Filter:  filter,
		Order:   sort.By + " " + sort.Direction,
	}
}
