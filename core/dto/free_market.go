package dto

import (
	"encoding/json"
	"io"
)

// FreeMarketRequestBody is an representation request body to create a new FreeMarket
type FreeMarketRequestBody struct {
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

// FromJSONFreeMarketRequestBody converts json body request to a FreeMarketRequestBody dto struct
func FromJSONFreeMarketRequestBody(body io.Reader) (*FreeMarketRequestBody, error) {
	FreeMarketRequest := FreeMarketRequestBody{}
	if err := json.NewDecoder(body).Decode(&FreeMarketRequest); err != nil {
		return nil, err
	}

	return &FreeMarketRequest, nil
}
