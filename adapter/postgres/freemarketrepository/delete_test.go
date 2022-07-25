package freemarketrepository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/booscaaa/go-api-test-unico/adapter/postgres/freemarketrepository"
	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/bxcodec/faker/v3"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func setupDelete() ([]string, domain.FreeMarket, sqlmock.Sqlmock, *sqlx.DB) {
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
	fakeFreeMarketResponse := domain.FreeMarket{}
	faker.FakeData(&fakeFreeMarketResponse)
	fakeFreeMarketResponse.AddressNumber = nil
	fakeFreeMarketResponse.Reference = nil

	mockDatabase, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDatabase, "sqlmock")

	return cols, fakeFreeMarketResponse, mock, sqlxDB
}

func TestDelete(t *testing.T) {
	t.Run("should delete a free market by id", func(t *testing.T) {
		cols, freeMarketResponse, mock, database := setupDelete()
		defer database.Close()

		mock.ExpectQuery("DELETE FROM free_market").WithArgs(
			freeMarketResponse.ID,
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
		freeMarket, err := sut.Delete(ctx, freeMarketResponse.ID)

		assert.Nil(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

		assert.Nil(t, err)
		assert.NotNil(t, freeMarketResponse)
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
	})

	t.Run("should return error on any database error", func(t *testing.T) {
		_, freeMarketResponse, mock, database := setupDelete()
		defer database.Close()

		mock.ExpectQuery("DELETE FROM free_market").WithArgs(
			freeMarketResponse.ID,
		).WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

		ctx := context.Background()
		sut := freemarketrepository.New(database)
		freeMarket, err := sut.Delete(ctx, freeMarketResponse.ID)

		assert.NotNil(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

		assert.Nil(t, freeMarket)
	})
}
