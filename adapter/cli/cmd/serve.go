/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"

	"github.com/booscaaa/go-api-test-unico/adapter/http/rest"
	"github.com/booscaaa/go-api-test-unico/adapter/postgres"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Initialize the REST api",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		postgres.RunMigrations()
		database := postgres.GetConnection(ctx)
		configHTTPServerPort := viper.GetString("server.http.port")
		rest.InitRest(configHTTPServerPort, database)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
