package freemarketservice_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/booscaaa/go-api-test-unico/adapter/http/rest/freemarketservice"
	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/booscaaa/go-api-test-unico/core/domain/mocks"
	"github.com/booscaaa/go-api-test-unico/util"
	"github.com/bxcodec/faker/v3"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zaptest"
)

func setupDelete(t *testing.T) (domain.FreeMarket, *gomock.Controller) {
	fakeFreeMarket := domain.FreeMarket{}
	faker.FakeData(&fakeFreeMarket)
	fakeFreeMarket.AddressNumber = nil
	fakeFreeMarket.Reference = nil

	mockCtrl := gomock.NewController(t)

	return fakeFreeMarket, mockCtrl
}

func TestDelete(t *testing.T) {
	t.Run("should delete a free market with status code 200", func(t *testing.T) {
		fakeFreeMarket, mock := setupDelete(t)
		defer mock.Finish()

		ctx := context.Background()
		validator := util.NewValidator()
		logger := zaptest.NewLogger(t)
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)
		mockFreeMarketUseCase.EXPECT().
			Delete(ctx, int64(1)).
			Return(&fakeFreeMarket, nil)

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator, logger)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market/1", nil)
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		sut.Delete(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return error on usecase exeption", func(t *testing.T) {
		_, mock := setupDelete(t)
		defer mock.Finish()

		ctx := context.Background()
		validator := util.NewValidator()
		logger := zaptest.NewLogger(t)
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)
		mockFreeMarketUseCase.EXPECT().
			Delete(ctx, int64(1)).
			Return(nil, fmt.Errorf("ANY ERROR"))

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator, logger)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market/1", nil)
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		sut.Delete(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 500 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return status code 500 when id not exists on request", func(t *testing.T) {
		_, mock := setupDelete(t)
		defer mock.Finish()

		validator := util.NewValidator()
		logger := zaptest.NewLogger(t)
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator, logger)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market", strings.NewReader("{"))
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		sut.Delete(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 500 {
			t.Errorf("status code is not correct")
		}
	})
}
