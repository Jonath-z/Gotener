package main

import (
	"github.com/joho/godotenv"
	"github.com/url_shortner/cmd"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Cannot load environment variables")
	}
	
	database := &cmd.Database{}
	db := database.Connect()
	database.InitializeModels(db)

	server := &cmd.Server{}
	server.InitializeRoutes()
	server.Run()
}
