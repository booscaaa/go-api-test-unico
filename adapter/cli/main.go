/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/booscaaa/go-api-test-unico/adapter/cli/cmd"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// @title UNICO API
// @description
// @contact.name Vinícius Boscardin
// @contact.email boscardinvinicius@gmail.com
// @BasePath /
func main() {
	cmd.Execute()
}
