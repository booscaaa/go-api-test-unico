package freemarketusecase

import (
	"context"

	"github.com/booscaaa/go-api-test-unico/core/domain"
)

func (usecase usecase) Delete(
	ctx context.Context,
	freeMarketID int64,
) (*domain.FreeMarket, error) {
	freeMarket, err := usecase.repository.Delete(ctx, freeMarketID)

	if err != nil {
		return nil, err
	}

	return freeMarket, nil
}
