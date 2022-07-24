package freemarketservice

import (
	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/go-playground/validator/v10"
)

type service struct {
	usecase   domain.FreeMarketUseCase
	validator validator.Validate
}

func New(
	usecase domain.FreeMarketUseCase,
	validator validator.Validate,
) domain.FreeMarketHTTPService {
	return &service{
		usecase:   usecase,
		validator: validator,
	}
}
