package freemarketrepository

import (
	"context"

	"github.com/booscaaa/go-api-test-unico/core/domain"
	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/booscaaa/go-paginate/paginate"
)

func (repository repository) Fetch(
	ctx context.Context,
	requestParams *dto.PaginationRequestParams,
) (*dto.PaginationResponseBody, error) {
	freeMarkets := []domain.FreeMarket{}
	count := int64(0)

	pagin := paginate.Instance(domain.FreeMarket{})

	pagin.Query("SELECT * FROM free_market").
		Page(requestParams.Page).
		Desc(requestParams.Descending).
		RowsPerPage(requestParams.ItemsPerPage).
		SearchBy(requestParams.Search, "region", "region_five", "market_name", "district").
		Sort(requestParams.Sort)

	query, queryCount := pagin.Select()

	err := repository.database.SelectContext(
		ctx, &freeMarkets, *query,
	)

	if err != nil {
		return nil, err
	}

	err = repository.database.QueryRowContext(
		ctx, *queryCount,
	).Scan(&count)

	if err != nil {
		return nil, err
	}

	return &dto.PaginationResponseBody{
		Items: freeMarkets,
		Total: count,
	}, nil
}
