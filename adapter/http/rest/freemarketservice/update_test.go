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
)

func TestUpdate(t *testing.T) {
	t.Run("should update a free market with status code 200", func(t *testing.T) {
		fakeFreeMarketRequest := dto.FreeMarketRequestBody{}
		fakeFreeMarket := domain.FreeMarket{}
		faker.FakeData(&fakeFreeMarketRequest)
		faker.FakeData(&fakeFreeMarket)

		mock := gomock.NewController(t)
		defer mock.Finish()

		ctx := context.Background()
		validator := util.NewValidator()
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)
		mockFreeMarketUseCase.EXPECT().
			Update(ctx, int64(1), &fakeFreeMarketRequest).
			Return(&fakeFreeMarket, nil)

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator)

		payload, _ := json.Marshal(fakeFreeMarketRequest)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market/1", strings.NewReader(string(payload)))
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		sut.Update(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 200 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return status code 500 when json is incorrect", func(t *testing.T) {
		fakeFreeMarketRequest := dto.FreeMarketRequestBody{}
		fakeFreeMarket := domain.FreeMarket{}
		faker.FakeData(&fakeFreeMarketRequest)
		faker.FakeData(&fakeFreeMarket)

		mock := gomock.NewController(t)
		defer mock.Finish()

		validator := util.NewValidator()
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market/1", strings.NewReader("{"))
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		sut.Update(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 500 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return status code 422 when json is missing requered key", func(t *testing.T) {
		fakeFreeMarketRequest := dto.FreeMarketRequestBody{}
		fakeFreeMarket := domain.FreeMarket{}
		faker.FakeData(&fakeFreeMarketRequest)
		faker.FakeData(&fakeFreeMarket)

		mock := gomock.NewController(t)
		defer mock.Finish()

		validator := util.NewValidator()
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market/1", strings.NewReader("{}"))
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		sut.Update(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 422 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return error on usecase exeption", func(t *testing.T) {
		fakeFreeMarketRequest := dto.FreeMarketRequestBody{}
		fakeFreeMarket := domain.FreeMarket{}
		faker.FakeData(&fakeFreeMarketRequest)
		faker.FakeData(&fakeFreeMarket)

		mock := gomock.NewController(t)
		defer mock.Finish()

		ctx := context.Background()
		validator := util.NewValidator()
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)
		mockFreeMarketUseCase.EXPECT().
			Update(ctx, int64(1), &fakeFreeMarketRequest).
			Return(nil, fmt.Errorf("ANY ERROR"))

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator)

		payload, _ := json.Marshal(fakeFreeMarketRequest)
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market/1", strings.NewReader(string(payload)))
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
		sut.Update(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 500 {
			t.Errorf("status code is not correct")
		}
	})

	t.Run("should return status code 500 when id not exists on request", func(t *testing.T) {
		fakeFreeMarketRequest := dto.FreeMarketRequestBody{}
		fakeFreeMarket := domain.FreeMarket{}
		faker.FakeData(&fakeFreeMarketRequest)
		faker.FakeData(&fakeFreeMarket)

		mock := gomock.NewController(t)
		defer mock.Finish()

		validator := util.NewValidator()
		mockFreeMarketUseCase := mocks.NewMockFreeMarketUseCase(mock)

		sut := freemarketservice.New(mockFreeMarketUseCase, *validator)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/free-market", strings.NewReader("{"))
		r.Header.Set("Content-Type", "application/json")

		c, _ := gin.CreateTestContext(w)
		c.Request = r
		sut.Update(c)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 500 {
			t.Errorf("status code is not correct")
		}
	})
}
