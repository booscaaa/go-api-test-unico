package freemarketservice

import (
	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type service struct {
	usecase   domain.FreeMarketUseCase
	validator validator.Validate
	logger    *zap.Logger
}

func New(
	usecase domain.FreeMarketUseCase,
	validator validator.Validate,
	logger *zap.Logger,
) domain.FreeMarketHTTPService {
	return &service{
		usecase:   usecase,
		validator: validator,
		logger:    logger,
	}
}
