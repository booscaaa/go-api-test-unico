package freemarketrepository

import (
	"context"

	"github.com/booscaaa/go-api-test-unico/core/domain"
)

func (repository repository) GetByID(
	ctx context.Context,
	freeMarketID int64,
) (*domain.FreeMarket, error) {
	freeMarket := domain.FreeMarket{}

	err := repository.database.GetContext(
		ctx, &freeMarket, "SELECT * FROM  free_market WHERE id = $1;", freeMarketID,
	)

	if err != nil {
		return nil, err
	}

	return &freeMarket, nil
}
