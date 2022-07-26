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
	Latitude          float64 `json:"latitude"          faker:"lat"         validate:"required"         example:"-28.460924"`
	Longitude         float64 `json:"longitude"         faker:"long"        validate:"required"         example:"-52.210348"`
	SetCens           string  `json:"setCens"           faker:"len=15"      validate:"required,max=15"  example:"355030885000091"`
	Areap             string  `json:"areaP"             faker:"len=13"      validate:"required,max=13"  example:"3550308005040"`
	RegionCode        int64   `json:"regionCode"        faker:"oneof: 1, 2" validate:"required"         example:"87"`
	Region            string  `json:"region"            faker:"len=18"      validate:"required,max=18"  example:"VILA FORMOSA"`
	SubprefectureCode int64   `json:"subprefectureCode" faker:"oneof: 1, 2" validate:"required"         example:"26"`
	Subprefecture     string  `json:"subprefecture"     faker:"len=25"      validate:"required,max=25"  example:"ARICANDUVA-FORMOSA-CARRAO"`
	RegionFive        string  `json:"regionFive"        faker:"len=6"       validate:"required,max=6"   example:"Leste"`
	RegionEight       string  `json:"regionEight"       faker:"len=7"       validate:"required,max=7"   example:"Leste 1"`
	MarketName        string  `json:"marketName"        faker:"len=30"      validate:"required,max=30"  example:"VILA FORMOSA"`
	Register          string  `json:"register"          faker:"len=6"       validate:"required,max=6"   example:"4041-0"`
	Address           string  `json:"address"           faker:"len=34"      validate:"required,max=34"  example:"RUA MARAGOJIPE"`
	AddressNumber     *string `json:"addressNumber"                         validate:"omitempty,lte=5"  example:"S/N"`
	District          string  `json:"district"          faker:"len=20"      validate:"required,max=20"  example:"VL FORMOSA"`
	Reference         *string `json:"reference"                             validate:"omitempty,lte=34" example:"TV RUA PRETORIA"`
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
