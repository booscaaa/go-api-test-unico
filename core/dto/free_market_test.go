package dto_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/bxcodec/faker/v3"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestFromJSONFreeMarketRequestBody(t *testing.T) {
	t.Run(
		"should convert the json request payload to a struct FreeMarketRequestBody",
		func(t *testing.T) {
			fakeItem := dto.FreeMarketRequestBody{}
			faker.FakeData(&fakeItem)
			fakeItem.AddressNumber = nil
			fakeItem.Reference = nil

			json, err := json.Marshal(fakeItem)
			assert.Nil(t, err)
			validate := validator.New()

			freeMarketRequest, err := dto.FromJSONFreeMarketRequestBody(
				strings.NewReader(string(json)),
				validate,
			)
			assert.Nil(t, err)
			assert.NotNil(t, freeMarketRequest)
			assert.Equal(t, fakeItem.Latitude, freeMarketRequest.Latitude)
			assert.Equal(t, fakeItem.Longitude, freeMarketRequest.Longitude)
			assert.Equal(t, fakeItem.SetCens, freeMarketRequest.SetCens)
			assert.Equal(t, fakeItem.Areap, freeMarketRequest.Areap)
			assert.Equal(t, fakeItem.RegionCode, freeMarketRequest.RegionCode)
			assert.Equal(t, fakeItem.Region, freeMarketRequest.Region)
			assert.Equal(t, fakeItem.SubprefectureCode, freeMarketRequest.SubprefectureCode)
			assert.Equal(t, fakeItem.Subprefecture, freeMarketRequest.Subprefecture)
			assert.Equal(t, fakeItem.RegionFive, freeMarketRequest.RegionFive)
			assert.Equal(t, fakeItem.RegionEight, freeMarketRequest.RegionEight)
			assert.Equal(t, fakeItem.MarketName, freeMarketRequest.MarketName)
			assert.Equal(t, fakeItem.Register, freeMarketRequest.Register)
			assert.Equal(t, fakeItem.Address, freeMarketRequest.Address)
			assert.Equal(t, fakeItem.AddressNumber, freeMarketRequest.AddressNumber)
			assert.Equal(t, fakeItem.District, freeMarketRequest.District)
			assert.Equal(t, fakeItem.Reference, freeMarketRequest.Reference)
		},
	)

	t.Run("should return error if json is incorrect", func(t *testing.T) {
		validate := validator.New()
		freeMarketRequest, err := dto.FromJSONFreeMarketRequestBody(strings.NewReader("{"), validate)

		assert.NotNil(t, err)
		assert.Nil(t, freeMarketRequest)
	})

	t.Run("should return error if json keys is incorrect", func(t *testing.T) {
		validate := validator.New()
		freeMarketRequest, err := dto.FromJSONFreeMarketRequestBody(strings.NewReader("{}"), validate)

		assert.NotNil(t, err)
		assert.Nil(t, freeMarketRequest)
	})
}
