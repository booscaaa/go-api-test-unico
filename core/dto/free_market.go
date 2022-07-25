package dto

import (
	"encoding/json"
	"io"

	"github.com/booscaaa/go-api-test-unico/util"
	"github.com/go-playground/validator/v10"
)

//ID,LONG,LAT,SETCENS,AREAP,CODDIST,DISTRITO,CODSUBPREF,SUBPREFE,REGIAO5,REGIAO8,NOME_FEIRA,REGISTRO,LOGRADOURO,NUMERO,BAIRRO,REFERENCIA

// FreeMarketRequestBody is an representation request body to create a new FreeMarket
type FreeMarketRequestBody struct {
	Latitude          float64 `json:"latitude"          faker:"lat"    validate:"required"`
	Longitude         float64 `json:"longitude"         faker:"long"   validate:"required"`
	SetCens           string  `json:"setCens"           faker:"len=15" validate:"required,max=15"`
	Areap             string  `json:"areaP"             faker:"len=13" validate:"required,max=13"`
	RegionCode        int64   `json:"regionCode"        faker:"oneof: 1, 2"   validate:"required"`
	Region            string  `json:"region"            faker:"len=18" validate:"required,max=18"`
	SubprefectureCode int64   `json:"subprefectureCode" faker:"oneof: 1, 2"   validate:"required"`
	Subprefecture     string  `json:"subprefecture"     faker:"len=25" validate:"required,max=25"`
	RegionFive        string  `json:"regionFive"        faker:"len=6"  validate:"required,max=6"`
	RegionEight       string  `json:"regionEight"       faker:"len=7"  validate:"required,max=7"`
	MarketName        string  `json:"marketName"        faker:"len=30" validate:"required,max=30"`
	Register          string  `json:"register"          faker:"len=6"  validate:"required,max=6"`
	Address           string  `json:"address"           faker:"len=34" validate:"required,max=34"`
	AddressNumber     *string `json:"addressNumber"`
	District          string  `json:"district"          faker:"len=20" validate:"required,max=20"`
	Reference         *string `json:"reference"`
}

// FromJSONFreeMarketRequestBody converts json body request to a FreeMarketRequestBody dto struct
func FromJSONFreeMarketRequestBody(
	body io.Reader,
	validate *validator.Validate,
) (*FreeMarketRequestBody, error) {
	freeMarketRequest := FreeMarketRequestBody{}
	if err := json.NewDecoder(body).Decode(&freeMarketRequest); err != nil {
		return nil, err
	}

	if err := validate.Struct(freeMarketRequest); err != nil {
		return nil, util.HandleValidatorFieldsError(err)
	}

	return &freeMarketRequest, nil
}
