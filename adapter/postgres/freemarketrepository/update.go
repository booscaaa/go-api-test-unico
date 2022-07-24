package freemarketrepository

import (
	"context"

	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/booscaaa/go-api-test-unico/core/dto"
)

func (repository repository) Update(
	ctx context.Context,
	freeMarketID int64,
	freeMarketRequest *dto.FreeMarketRequestBody,
) (*domain.FreeMarket, error) {
	freeMarket := domain.FreeMarket{}

	err := repository.database.QueryRowxContext(
		ctx,
		`UPDATE free_market 
			SET latitude = $1,
			SET longitude = $2,
			SET set_cens = $3,
			SET areap = $4,
			SET region_code = $5,
			SET region = $6,
			SET subprefecture_code = $7,
			SET subprefecture = $8,
			SET region_five = $9,
			SET region_eight = $10,
			SET market_name = $11,
			SET register = $12,
			SET address = $13,
			SET address_number = $14,
			SET district = $15,
			SET reference = $16
		WHERE id = $17 returning *;`,
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
		freeMarketID,
	).StructScan(&freeMarket)

	if err != nil {
		return nil, err
	}

	return &freeMarket, nil
}
