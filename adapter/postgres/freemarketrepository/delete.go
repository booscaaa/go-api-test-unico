package freemarketrepository

import (
	"context"

	"github.com/booscaaa/go-api-test-unico/core/domain"
)

func (repository repository) Delete(
	ctx context.Context,
	freeMarketID int64,
) (*domain.FreeMarket, error) {
	freeMarket := domain.FreeMarket{}

	err := repository.database.QueryRowxContext(
		ctx, "DELETE FROM  free_market WHERE id = $1;", freeMarketID,
	).StructScan(&freeMarket)

	if err != nil {
		return nil, err
	}

	return &freeMarket, nil
}
