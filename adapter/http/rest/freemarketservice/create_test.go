package freemarketservice_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
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

func setupCreate(t *testing.T) (dto.FreeMarketRequestBody, domain.FreeMarket, *gomock.Controller) {
	fakeFreeMarketRequest := dto.FreeMarketRequestBody{}
	fakeFreeMarket := domain.FreeMarket{}
	faker.FakeData(&fakeFreeMarketRequest)
	faker.FakeData(&fakeFreeMarket)

	mockCtrl := gomock.NewController(t)

	return fakeFreeMarketRequest, fakeFreeMarket, mockCtrl
}

func TestCreate(t *testing.T) {
	t.Run("should create a free market with status code 200", func(t *testing.T) {
		fakeFreeMarketRequest, fakeFreeMarket, mock := setupCreate(t)
		defer mock.Finish()

		ctx := context.Background()
		validator := util.NewValidator()
		logger := zaptest.NewLogger(t)
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)
		mockFreeMarketUseCase.EXPECT().
			Create(ctx, &fakeFreeMarketRequest).
			Return(&fakeFreeMarket, nil)

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator, logger)

		payload, _ := json.Marshal(fakeFreeMarketRequest)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market", strings.NewReader(string(payload)))
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		sut.Create(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return status code 500 when json is incorrect", func(t *testing.T) {
		_, _, mock := setupCreate(t)
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
		sut.Create(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 500 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return status code 422 when json is missing requered key", func(t *testing.T) {
		_, _, mock := setupCreate(t)
		defer mock.Finish()

		validator := util.NewValidator()
		logger := zaptest.NewLogger(t)
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator, logger)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market", strings.NewReader("{}"))
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		sut.Create(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 422 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return error on usecase exeption", func(t *testing.T) {
		fakeFreeMarketRequest, _, mock := setupCreate(t)
		defer mock.Finish()

		ctx := context.Background()
		validator := util.NewValidator()
		logger := zaptest.NewLogger(t)

		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)
		mockFreeMarketUseCase.EXPECT().
			Create(ctx, &fakeFreeMarketRequest).
			Return(nil, fmt.Errorf("ANY ERROR"))

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator, logger)

		payload, _ := json.Marshal(fakeFreeMarketRequest)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market", strings.NewReader(string(payload)))
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		sut.Create(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 500 {
			t.Errorf("status code is not correct")
		}
	})
}
