package domain

import (
	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/gin-gonic/gin"
)

type FreeMarket struct {
	ID                int64   `json:"id"                db:"id"`
	Latitude          float64 `json:"latitude"          db:"latitude"`
	Longitude         float64 `json:"longitude"         db:"longitude"`
	SetCens           string  `json:"setCens"           db:"set_cens"`
	Areap             string  `json:"areaP"             db:"areap"`
	RegionCode        int64   `json:"regionCode"        db:"region_code"`
	Region            string  `json:"region"            db:"region"`
	SubprefectureCode int64   `json:"subprefectureCode" db:"subprefecture_code"`
	Subprefecture     string  `json:"subprefecture"     db:"subprefecture"`
	RegionFive        string  `json:"regionFive"        db:"region_five"`
	RegionEight       string  `json:"regionEight"       db:"region_eight"`
	MarketName        string  `json:"marketName"        db:"market_name"`
	Register          string  `json:"register"          db:"register"`
	Address           string  `json:"address"           db:"address"`
	AddressNumber     *string `json:"addressNumber"     db:"address_number"`
	District          string  `json:"district"          db:"disctrict"`
	Reference         *string `json:"reference"         db:"reference"`
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
