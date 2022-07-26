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
		`UPDATE free_market SET 
			latitude = $1,
			longitude = $2,
			set_cens = $3,
			areap = $4,
			region_code = $5,
			region = $6,
			subprefecture_code = $7,
			subprefecture = $8,
			region_five = $9,
			region_eight = $10,
			market_name = $11,
			register = $12,
			address = $13,
			address_number = $14,
			district = $15,
			reference = $16
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
