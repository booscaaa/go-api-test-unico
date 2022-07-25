package freemarketservice_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/booscaaa/go-api-test-unico/adapter/http/rest/freemarketservice"
	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/booscaaa/go-api-test-unico/core/domain/mocks"
	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/booscaaa/go-api-test-unico/util"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zaptest"
)

func setupFetch(t *testing.T) (domain.FreeMarket, *gomock.Controller) {
	fakeFreeMarket := domain.FreeMarket{}
	faker.FakeData(&fakeFreeMarket)
	fakeFreeMarket.AddressNumber = nil
	fakeFreeMarket.Reference = nil

	mockCtrl := gomock.NewController(t)

	return fakeFreeMarket, mockCtrl
}

func TestFetch(t *testing.T) {
	t.Run("should fetch a list of free market with status code 200", func(t *testing.T) {
		fakeFreeMarket, mock := setupFetch(t)
		defer mock.Finish()

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market", nil)
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)

		queryStringParams := r.URL.Query()
		queryStringParams.Add("page", "1")
		queryStringParams.Add("itemsPerPage", "11")
		queryStringParams.Add("sort", "")
		queryStringParams.Add("descending", "")
		queryStringParams.Add("search", "")
		queryStringParams.Add("status", "0")
		r.URL.RawQuery = queryStringParams.Encode()

		c.Request = r

		requestParams := dto.FromValuePaginationRequestParams(c.Request)

		ctx := context.Background()
		validator := util.NewValidator()
		logger := zaptest.NewLogger(t)
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)
		mockFreeMarketUseCase.EXPECT().
			Fetch(ctx, requestParams).
			Return(&dto.PaginationResponseBody{
				Items: []domain.FreeMarket{fakeFreeMarket},
				Total: 1,
			}, nil)

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator, logger)

		sut.Fetch(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return error on usecase exeption", func(t *testing.T) {
		_, mock := setupFetch(t)
		defer mock.Finish()

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market/1", nil)
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		queryStringParams := r.URL.Query()
		queryStringParams.Add("page", "1")
		queryStringParams.Add("itemsPerPage", "11")
		queryStringParams.Add("sort", "")
		queryStringParams.Add("descending", "")
		queryStringParams.Add("search", "")
		queryStringParams.Add("status", "0")
		r.URL.RawQuery = queryStringParams.Encode()
		c.Request = r

		requestParams := dto.FromValuePaginationRequestParams(c.Request)

		ctx := context.Background()
		validator := util.NewValidator()
		logger := zaptest.NewLogger(t)
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)
		mockFreeMarketUseCase.EXPECT().
			Fetch(ctx, requestParams).
			Return(nil, fmt.Errorf("ANY ERROR"))

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator, logger)

		sut.Fetch(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 500 {
			t.Errorf("status code is not correct")
		}
	})

}
