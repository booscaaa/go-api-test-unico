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
	Latitude          float64 `json:"latitude"          faker:"lat"    validate:"required"        csv:"LAT"`
	Longitude         float64 `json:"longitude"         faker:"long"   validate:"required"        csv:"LONG"`
	SetCens           string  `json:"setCens"           faker:"len=15" validate:"required,max=15" csv:"SETCENS"`
	Areap             string  `json:"areaP"             faker:"len=13" validate:"required,max=13" csv:"AREAP"`
	RegionCode        int64   `json:"regionCode"                       validate:"required"        csv:"CODDIST"`
	Region            string  `json:"region"            faker:"len=18" validate:"required,max=18" csv:"DISTRITO"`
	SubprefectureCode int64   `json:"subprefectureCode"                validate:"required"        csv:"CODSUBPREF"`
	Subprefecture     string  `json:"subprefecture"     faker:"len=25" validate:"required,max=25" csv:"SUBPREFE"`
	RegionFive        string  `json:"regionFive"        faker:"len=6"  validate:"required,max=6"  csv:"REGIAO5"`
	RegionEight       string  `json:"regionEight"       faker:"len=7"  validate:"required,max=7"  csv:"REGIAO8"`
	MarketName        string  `json:"marketName"        faker:"len=30" validate:"required,max=30" csv:"NOME_FEIRA"`
	Register          string  `json:"register"          faker:"len=6"  validate:"required,max=6"  csv:"REGISTRO"`
	Address           string  `json:"address"           faker:"len=34" validate:"required,max=34" csv:"LOGRADOURO"`
	AddressNumber     *string `json:"addressNumber"                                               csv:"NUMERO"`
	District          string  `json:"district"          faker:"len=20" validate:"required,max=20" csv:"BAIRRO"`
	Reference         *string `json:"reference"                                                   csv:"REFERENCIA"`
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
