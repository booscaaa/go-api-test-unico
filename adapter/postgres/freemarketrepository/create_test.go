package freemarketrepository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/booscaaa/go-api-test-unico/adapter/postgres/freemarketrepository"
	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/bxcodec/faker/v3"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func setupCreate() ([]string, dto.FreeMarketRequestBody, domain.FreeMarket, sqlmock.Sqlmock, *sqlx.DB) {
	cols := []string{
		"id",
		"latitude",
		"longitude",
		"set_cens",
		"areap",
		"region_code",
		"region",
		"subprefecture_code",
		"subprefecture",
		"region_five",
		"region_eight",
		"market_name",
		"register",
		"address",
		"address_number",
		"district",
		"reference",
	}
	fakeFreeMarketRequest := dto.FreeMarketRequestBody{}
	fakeFreeMarketResponse := domain.FreeMarket{}
	faker.FakeData(&fakeFreeMarketRequest)
	faker.FakeData(&fakeFreeMarketResponse)

	mockDatabase, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDatabase, "sqlmock")

	return cols, fakeFreeMarketRequest, fakeFreeMarketResponse, mock, sqlxDB
}

func TestCreate(t *testing.T) {
	t.Run("should create a free market", func(t *testing.T) {
		cols, freeMarketRequest, freeMarketResponse, mock, database := setupCreate()
		defer database.Close()

		mock.ExpectQuery("INSERT INTO free_market (.+)").WithArgs(
			freeMarketRequest.Latitude,
			freeMarketRequest.Longitude,
			freeMarketRequest.SetCens,
			freeMarketRequest.Areap,
			freeMarketRequest.RegionCode,
			freeMarketRequest.Region,
			freeMarketRequest.SubprefectureCode,
			freeMarketRequest.Subprefecture,
			freeMarketRequest.RegionFive,
			freeMarketRequest.RegionEight,
			freeMarketRequest.MarketName,
			freeMarketRequest.Register,
			freeMarketRequest.Address,
			freeMarketRequest.AddressNumber,
			freeMarketRequest.District,
			freeMarketRequest.Reference,
		).WillReturnRows(sqlmock.NewRows(cols).AddRow(
			freeMarketResponse.ID,
			freeMarketResponse.Latitude,
			freeMarketResponse.Longitude,
			freeMarketResponse.SetCens,
			freeMarketResponse.Areap,
			freeMarketResponse.RegionCode,
			freeMarketResponse.Region,
			freeMarketResponse.SubprefectureCode,
			freeMarketResponse.Subprefecture,
			freeMarketResponse.RegionFive,
			freeMarketResponse.RegionEight,
			freeMarketResponse.MarketName,
			freeMarketResponse.Register,
			freeMarketResponse.Address,
			freeMarketResponse.AddressNumber,
			freeMarketResponse.District,
			freeMarketResponse.Reference,
		))

		ctx := context.Background()
		sut := freemarketrepository.New(database)
		freeMarket, err := sut.Create(ctx, &freeMarketRequest)

		assert.Nil(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

		assert.Nil(t, err)
		assert.NotNil(t, freeMarketResponse)
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

	t.Run("should return error on any database error", func(t *testing.T) {
		_, freeMarketRequest, _, mock, database := setupCreate()
		defer database.Close()

		mock.ExpectQuery("INSERT INTO free_market (.+)").WithArgs(
			freeMarketRequest.Latitude,
			freeMarketRequest.Longitude,
			freeMarketRequest.SetCens,
			freeMarketRequest.Areap,
			freeMarketRequest.RegionCode,
			freeMarketRequest.Region,
			freeMarketRequest.SubprefectureCode,
			freeMarketRequest.Subprefecture,
			freeMarketRequest.RegionFive,
			freeMarketRequest.RegionEight,
			freeMarketRequest.MarketName,
			freeMarketRequest.Register,
			freeMarketRequest.Address,
			freeMarketRequest.AddressNumber,
			freeMarketRequest.District,
			freeMarketRequest.Reference,
		).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

		ctx := context.Background()
		sut := freemarketrepository.New(database)
		freeMarket, err := sut.Create(ctx, &freeMarketRequest)

		assert.NotNil(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		assert.Nil(t, freeMarket)
	})
}
