package domain

import (
	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/gin-gonic/gin"
)

type FreeMarket struct {
	ID                int64   `json:"id"`
	Latitude          float64 `json:"latitude"`
	Longitude         float64 `json:"longitude"`
	SetCens           string  `json:"setCens"`
	Areap             string  `json:"areaP"`
	RegionCode        int64   `json:"regionCode"`
	Region            string  `json:"region"`
	SubprefectureCode int64   `json:"subprefectureCode"`
	Subprefecture     string  `json:"subprefecture"`
	RegionFive        string  `json:"regionFive"`
	RegionEight       string  `json:"regionEight"`
	MarketName        string  `json:"marketName"`
	Register          string  `json:"register"`
	Adress            string  `json:"adress"`
	AdressNumber      *string `json:"adressNumber"`
	District          string  `json:"district"`
	Reference         *string `json:"reference"`
}

// FreeMarketService is a contract of http adapter layer
type FreeMarketHTTPService interface {
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
	GetByID(*gin.Context)
	Fetch(*gin.Context)
}

// FreeMarketUseCase is a contract of usecase ruler layer
type FreeMarketUseCase interface {
	Create(*dto.FreeMarketRequestBody) (*FreeMarket, error)
	Update(int, *dto.FreeMarketRequestBody) (*FreeMarket, error)
	Delete(int) (*FreeMarket, error)
	GetByID(int) (*FreeMarket, error)
	Fetch(*dto.PaginationRequestParms) (*dto.PaginationResponseBody, error)
}

// FreeMarketRepository is a contract of database connection adapter layer
type FreeMarketRepository interface {
	Create(*dto.FreeMarketRequestBody) (*FreeMarket, error)
	Update(int, *dto.FreeMarketRequestBody) (*FreeMarket, error)
	Delete(int) (*FreeMarket, error)
	GetByID(int) (*FreeMarket, error)
	Fetch(*dto.PaginationRequestParms) (*dto.PaginationResponseBody, error)
}
