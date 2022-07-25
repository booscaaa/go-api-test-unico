package freemarketusecase

import "github.com/booscaaa/go-api-test-unico/core/domain"

type usecase struct {
	repository domain.FreeMarketRepository
}

func New(repository domain.FreeMarketRepository) domain.FreeMarketUseCase {
	return &usecase{
		repository: repository,
	}
}
