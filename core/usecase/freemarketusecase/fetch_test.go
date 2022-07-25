package freemarketusecase_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/booscaaa/go-api-test-unico/core/domain/mocks"
	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/booscaaa/go-api-test-unico/core/usecase/freemarketusecase"
	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	t.Run("should get a free market pagination by pagination params", func(t *testing.T) {
		fakeFreeMarketRequest := dto.FreeMarketRequestBody{}
		freeMarketResponse := domain.FreeMarket{}
		faker.FakeData(&fakeFreeMarketRequest)
		faker.FakeData(&freeMarketResponse)

		requestParams := dto.PaginationRequestParams{
			Search:       "",
			Descending:   []string{},
			Sort:         []string{},
			Page:         1,
			ItemsPerPage: 10,
			Status:       0,
		}

		freeMarketResponsePaginated := dto.PaginationResponseBody{
			Items: []domain.FreeMarket{freeMarketResponse},
			Total: 1,
		}

		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockFreeMarketRepository := mocks.NewMockFreeMarketRepository(mockCtrl)
		mockFreeMarketRepository.EXPECT().
			Fetch(ctx, &requestParams).
			Return(&freeMarketResponsePaginated, nil)

		sut := freemarketusecase.New(mockFreeMarketRepository)
		freeMarketsPaginated, err := sut.Fetch(ctx, &requestParams)

		assert.Nil(t, err)
		for _, freeMarket := range freeMarketsPaginated.Items.([]domain.FreeMarket) {
			assert.Equal(t, freeMarket.ID, freeMarketResponse.ID)
			assert.Equal(t, freeMarket.Latitude, freeMarketResponse.Latitude)
			assert.Equal(t, freeMarket.Latitude, freeMarketResponse.Latitude)
			assert.Equal(t, freeMarket.Longitude, freeMarketResponse.Longitude)
			assert.Equal(t, freeMarket.SetCens, freeMarketResponse.SetCens)
			assert.Equal(t, freeMarket.Areap, freeMarketResponse.Areap)
			assert.Equal(t, freeMarket.RegionCode, freeMarketResponse.RegionCode)
			assert.Equal(t, freeMarket.Region, freeMarketResponse.Region)
			assert.Equal(t, freeMarket.SubprefectureCode, freeMarketResponse.SubprefectureCode)
			assert.Equal(t, freeMarket.Subprefecture, freeMarketResponse.Subprefecture)
			assert.Equal(t, freeMarket.RegionFive, freeMarketResponse.RegionFive)
			assert.Equal(t, freeMarket.RegionEight, freeMarketResponse.RegionEight)
			assert.Equal(t, freeMarket.MarketName, freeMarketResponse.MarketName)
			assert.Equal(t, freeMarket.Register, freeMarketResponse.Register)
			assert.Equal(t, freeMarket.Address, freeMarketResponse.Address)
			assert.Equal(t, freeMarket.AddressNumber, freeMarketResponse.AddressNumber)
			assert.Equal(t, freeMarket.District, freeMarketResponse.District)
			assert.Equal(t, freeMarket.Reference, freeMarketResponse.Reference)
		}
	})

	t.Run("should return error on get a free market paginated", func(t *testing.T) {
		fakeFreeMarketRequest := dto.FreeMarketRequestBody{}
		freeMarketResponse := domain.FreeMarket{}
		faker.FakeData(&fakeFreeMarketRequest)
		faker.FakeData(&freeMarketResponse)

		requestParams := dto.PaginationRequestParams{
			Search:       "",
			Descending:   []string{},
			Sort:         []string{},
			Page:         1,
			ItemsPerPage: 10,
			Status:       0,
		}

		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockFreeMarketRepository := mocks.NewMockFreeMarketRepository(mockCtrl)
		mockFreeMarketRepository.EXPECT().
			Fetch(ctx, &requestParams).
			Return(nil, fmt.Errorf("ANY ERROR"))

		sut := freemarketusecase.New(mockFreeMarketRepository)
		freeMarket, err := sut.Fetch(ctx, &requestParams)

		assert.NotNil(t, err)
		assert.Nil(t, freeMarket)
	})
}
