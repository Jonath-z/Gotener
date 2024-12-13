package main

import (
	"github.com/joho/godotenv"
	"github.com/url_shortner/cmd"
	"github.com/url_shortner/db"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Cannot load environment variables")
	}

	database := &db.Database{}
	database.Connect()
	cmd.InitializeModels(database.Db)

	server := &cmd.Server{}
	server.InitializeRoutes()
	server.Run()
}
