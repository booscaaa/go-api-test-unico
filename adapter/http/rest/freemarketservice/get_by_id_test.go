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
)

func setupGetByID(t *testing.T) (domain.FreeMarket, *gomock.Controller) {
	fakeFreeMarket := domain.FreeMarket{}
	faker.FakeData(&fakeFreeMarket)

	mockCtrl := gomock.NewController(t)

	return fakeFreeMarket, mockCtrl
}

func TestGetByID(t *testing.T) {
	t.Run("should get a free market by id with status code 200", func(t *testing.T) {
		fakeFreeMarket, mock := setupGetByID(t)
		defer mock.Finish()

		ctx := context.Background()
		validator := util.NewValidator()
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)
		mockFreeMarketUseCase.EXPECT().
			GetByID(ctx, int64(1)).
			Return(&fakeFreeMarket, nil)

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market/1", nil)
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		sut.GetByID(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return error on usecase exeption", func(t *testing.T) {
		_, mock := setupGetByID(t)
		defer mock.Finish()

		ctx := context.Background()
		validator := util.NewValidator()
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)
		mockFreeMarketUseCase.EXPECT().
			GetByID(ctx, int64(1)).
			Return(nil, fmt.Errorf("ANY ERROR"))

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market/1", nil)
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		sut.GetByID(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 500 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return status code 500 when id not exists on request", func(t *testing.T) {
		_, mock := setupGetByID(t)
		defer mock.Finish()

		validator := util.NewValidator()
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market", strings.NewReader("{"))
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		sut.GetByID(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 500 {
			t.Errorf("status code is not correct")
		}
	})
}
