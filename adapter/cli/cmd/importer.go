/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/booscaaa/go-api-test-unico/adapter/postgres"
	"github.com/booscaaa/go-api-test-unico/adapter/postgres/freemarketrepository"
	"github.com/booscaaa/go-api-test-unico/core/dto"
	"github.com/booscaaa/go-api-test-unico/core/usecase/freemarketusecase"
	"github.com/spf13/cobra"
)

// importerCmd represents the importer command
var importerCmd = &cobra.Command{
	Use:   "importer",
	Short: "Import the CSV file into database",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		postgres.RunMigrations()
		database := postgres.GetConnection(ctx)
		freeMarketRepository := freemarketrepository.New(database)
		freeMarketUsecase := freemarketusecase.New(freeMarketRepository)

		records, err := readCsvData("data-import/DEINFO_AB_FEIRASLIVRES_2014.csv")

		if err != nil {
			log.Println(err)
			return
		}

		for _, record := range records {
			latitude, err := strconv.ParseFloat(record[2], 64)
			if err != nil {
				log.Println(err)
				return
			}

			longitude, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				log.Println(err)
				return
			}

			codDist, err := strconv.Atoi(record[5])
			if err != nil {
				log.Println(err)
				return
			}

			subprefectureCode, err := strconv.Atoi(record[7])
			if err != nil {
				log.Println(err)
				return
			}

			freeMarket := dto.FreeMarketRequestBody{
				Latitude:          changeLatLongFloatDotPosition(latitude),
				Longitude:         changeLatLongFloatDotPosition(longitude),
				SetCens:           record[3],
				Areap:             record[4],
				RegionCode:        int64(codDist),
				Region:            record[6],
				SubprefectureCode: int64(subprefectureCode),
				Subprefecture:     record[8],
				RegionFive:        record[9],
				RegionEight:       record[10],
				MarketName:        record[11],
				Register:          record[12],
				Address:           record[13],
				AddressNumber:     &strings.Split(record[14], ".")[0],
				District:          record[15],
				Reference:         &record[16],
			}

			_, err = freeMarketUsecase.Create(ctx, &freeMarket)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(freeMarket.Latitude)
		}
	},
}

func changeLatLongFloatDotPosition(value float64) float64 {
	return value / 1000000
}

func readCsvData(fileName string) ([][]string, error) {
	f, err := os.Open(fileName)

	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	r := csv.NewReader(f)

	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	records, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}

func init() {
	rootCmd.AddCommand(importerCmd)
}
