package freemarketrepository

import (
	"context"

	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/booscaaa/go-api-test-unico/core/dto"
)

func (repository repository) Create(
	ctx context.Context,
	freeMarketRequest *dto.FreeMarketRequestBody,
) (*domain.FreeMarket, error) {
	freeMarket := domain.FreeMarket{}

	err := repository.database.QueryRowxContext(
		ctx,
		`INSERT INTO free_market (
			latitude,
			longitude,
			set_cens,
			areap,
			region_code,
			region,
			subprefecture_code,
			subprefecture,
			region_five,
			region_eight,
			market_name,
			register,
			address,
			address_number,
			district,
			reference
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`,
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
	).StructScan(&freeMarket)

	if err != nil {
		return nil, err
	}

	return &freeMarket, nil
}
