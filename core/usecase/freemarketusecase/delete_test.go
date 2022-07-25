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

func TestDelete(t *testing.T) {
	t.Run("should delete a free market by free market id", func(t *testing.T) {
		fakeFreeMarketRequest := dto.FreeMarketRequestBody{}
		freeMarketResponse := domain.FreeMarket{}
		faker.FakeData(&fakeFreeMarketRequest)
		faker.FakeData(&freeMarketResponse)

		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockFreeMarketRepository := mocks.NewMockFreeMarketRepository(mockCtrl)
		mockFreeMarketRepository.EXPECT().
			Delete(ctx, freeMarketResponse.ID).
			Return(&freeMarketResponse, nil)

		sut := freemarketusecase.New(mockFreeMarketRepository)
		freeMarket, err := sut.Delete(ctx, freeMarketResponse.ID)

		assert.Nil(t, err)
		assert.NotNil(t, freeMarket)
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
	})

	t.Run("should return error on delete free market", func(t *testing.T) {
		fakeFreeMarketRequest := dto.FreeMarketRequestBody{}
		freeMarketResponse := domain.FreeMarket{}
		faker.FakeData(&fakeFreeMarketRequest)
		faker.FakeData(&freeMarketResponse)

		ctx := context.Background()
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		mockFreeMarketRepository := mocks.NewMockFreeMarketRepository(mockCtrl)
		mockFreeMarketRepository.EXPECT().
			Delete(ctx, freeMarketResponse.ID).
			Return(nil, fmt.Errorf("ANY ERROR"))

		sut := freemarketusecase.New(mockFreeMarketRepository)
		freeMarket, err := sut.Delete(ctx, freeMarketResponse.ID)

		assert.NotNil(t, err)
		assert.Nil(t, freeMarket)
	})
}
