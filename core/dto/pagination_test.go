package dto_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/stretchr/testify/assert"
)

func TestFromValuePaginationRequestParams(t *testing.T) {
	t.Run(
		"should convert the query string parameters to a struct PaginationRequestParms",
		func(t *testing.T) {
			fakeRequest := httptest.NewRequest(http.MethodGet, "/free-market", nil)
			queryStringParams := fakeRequest.URL.Query()
			queryStringParams.Add("page", "1")
			queryStringParams.Add("itemsPerPage", "10")
			queryStringParams.Add("sort", "")
			queryStringParams.Add("descending", "")
			queryStringParams.Add("search", "")
			fakeRequest.URL.RawQuery = queryStringParams.Encode()

			paginationRequest := dto.FromValuePaginationRequestParams(fakeRequest)

			assert.Equal(t, paginationRequest.Page, 1)
			assert.Equal(t, paginationRequest.ItemsPerPage, 10)
			assert.Equal(t, paginationRequest.Sort, []string{""})
			assert.Equal(t, paginationRequest.Descending, []string{""})
			assert.Equal(t, paginationRequest.Search, "")
		},
	)
}
