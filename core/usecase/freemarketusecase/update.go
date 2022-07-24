package freemarketusecase

import (
	"context"

	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/booscaaa/go-api-test-unico/core/dto"
)

func (usecase usecase) Update(
	ctx context.Context,
	freeMarketID int64,
	freeMarketRequest *dto.FreeMarketRequestBody,
) (*domain.FreeMarket, error) {
	freeMarket, err := usecase.repository.Update(ctx, freeMarketID, freeMarketRequest)

	if err != nil {
		return nil, err
	}

	return freeMarket, nil
}
