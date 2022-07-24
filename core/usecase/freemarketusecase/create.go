package freemarketusecase

import (
	"context"

	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/booscaaa/go-api-test-unico/core/dto"
)

func (usecase usecase) Create(
	ctx context.Context,
	freeMarketRequest *dto.FreeMarketRequestBody,
) (*domain.FreeMarket, error) {
	freeMarket, err := usecase.repository.Create(ctx, freeMarketRequest)

	if err != nil {
		return nil, err
	}

	return freeMarket, nil
}
