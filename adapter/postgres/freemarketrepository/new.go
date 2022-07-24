package freemarketrepository

import (
	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	database *sqlx.DB
}

func New(database *sqlx.DB) domain.FreeMarketRepository {
	return &repository{
		database: database,
	}
}
