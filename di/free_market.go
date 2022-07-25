package di

import (
	"github.com/booscaaa/go-api-test-unico/adapter/http/rest/freemarketservice"
	"github.com/booscaaa/go-api-test-unico/adapter/postgres/freemarketrepository"
	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/booscaaa/go-api-test-unico/core/usecase/freemarketusecase"
	"github.com/booscaaa/go-api-test-unico/util"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

func ConfigFreeMarketDi(database *sqlx.DB, logger *zap.Logger) domain.FreeMarketHTTPService {
	validator := util.NewValidator()
	freeMarketRepository := freemarketrepository.New(database)
	freeMarketUsecase := freemarketusecase.New(freeMarketRepository)
	freeMarketService := freemarketservice.New(freeMarketUsecase, *validator, logger)

	return freeMarketService
}
