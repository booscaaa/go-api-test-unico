package freemarketusecase

import (
	"context"

	"github.com/booscaaa/go-api-test-unico/core/dto"
)

func (usecase usecase) Fetch(
	ctx context.Context,
	requestParams *dto.PaginationRequestParams,
) (*dto.PaginationResponseBody, error) {
	freeMarket, err := usecase.repository.Fetch(ctx, requestParams)

	if err != nil {
		return nil, err
	}

	return freeMarket, nil
}
