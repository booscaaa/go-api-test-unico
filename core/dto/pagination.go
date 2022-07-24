package dto

import (
	"net/http"
	"strconv"
	"strings"
)

// PaginationRequestParms is an representation query string params to filter and paginate data
type PaginationRequestParams struct {
	Search       string   `json:"search"`
	Descending   []string `json:"descending"`
	Page         int      `json:"page"`
	ItemsPerPage int      `json:"itemsPerPage"`
	Sort         []string `json:"sort"`
	Status       int      `json:"status"`
}

// PaginationResponseBody is an representation to paginated data
type PaginationResponseBody struct {
	Items interface{} `json:"items"`
	Total int64       `json:"total"`
}

// FromValuePaginationRequestParams converts query string params to a PaginationRequestParms struct
func FromValuePaginationRequestParams(request *http.Request) *PaginationRequestParams {
	page, _ := strconv.Atoi(request.FormValue("page"))
	itemsPerPage, _ := strconv.Atoi(request.FormValue("itemsPerPage"))
	status, _ := strconv.Atoi(request.FormValue("status"))

	paginationRequestParams := PaginationRequestParams{
		Search:       request.FormValue("search"),
		Descending:   strings.Split(request.FormValue("descending"), ","),
		Sort:         strings.Split(request.FormValue("sort"), ","),
		Page:         page,
		ItemsPerPage: itemsPerPage,
		Status:       status,
	}

	return &paginationRequestParams
}
