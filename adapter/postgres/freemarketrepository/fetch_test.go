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

func setupFetch() ([]string, dto.PaginationRequestParams, domain.FreeMarket, sqlmock.Sqlmock, *sqlx.DB) {
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
	requestParams := dto.PaginationRequestParams{
		Search:       "",
		Descending:   []string{},
		Sort:         []string{},
		Page:         1,
		ItemsPerPage: 10,
		Status:       0,
	}

	mockDatabase, mock, _ := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDatabase, "sqlmock")

	return cols, requestParams, fakeFreeMarketResponse, mock, sqlxDB
}

func TestFetch(t *testing.T) {
	t.Run("should return a list of free market by pagination params", func(t *testing.T) {
		cols, requestParams, freeMarketResponse, mock, database := setupFetch()
		defer database.Close()

		mock.ExpectQuery("SELECT (.+) FROM free_market").
			WillReturnRows(sqlmock.NewRows(cols).AddRow(
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

		mock.ExpectQuery("SELECT COUNT(.+) FROM free_market").
			WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(int64(requestParams.ItemsPerPage)))

		ctx := context.Background()
		sut := freemarketrepository.New(database)
		freeMarketsPaginated, err := sut.Fetch(ctx, &requestParams)

		assert.Nil(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

		assert.Nil(t, err)
		assert.NotNil(t, freeMarketsPaginated)

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

	t.Run("should return error on any database error into select fields query", func(t *testing.T) {
		_, requestParams, _, mock, database := setupFetch()
		defer database.Close()

		mock.ExpectQuery("SELECT (.+) FROM free_market").
			WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

		ctx := context.Background()
		sut := freemarketrepository.New(database)
		freeMarket, err := sut.Fetch(ctx, &requestParams)

		assert.NotNil(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

		assert.Nil(t, freeMarket)
	})

	t.Run("should return error on any database error into select count query", func(t *testing.T) {
		cols, requestParams, freeMarketResponse, mock, database := setupFetch()
		defer database.Close()

		mock.ExpectQuery("SELECT (.+) FROM free_market").
			WillReturnRows(sqlmock.NewRows(cols).AddRow(
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

		mock.ExpectQuery("SELECT COUNT(.+) FROM free_market").
			WillReturnError(fmt.Errorf("ANY DATABASE ERROR"))

		ctx := context.Background()
		sut := freemarketrepository.New(database)
		freeMarket, err := sut.Fetch(ctx, &requestParams)

		assert.NotNil(t, err)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

		assert.Nil(t, freeMarket)
	})
}
