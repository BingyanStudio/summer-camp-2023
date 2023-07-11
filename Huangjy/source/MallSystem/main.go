package main

import (
	"MallSystem/database"
	"MallSystem/routes"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("config/config.json")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	database.InitDatabase()

	e := routes.SetupRoutes()

	e.Run(":8080")
}
