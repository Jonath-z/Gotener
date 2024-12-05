package cmd

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/url_shortner/lib"
	"github.com/url_shortner/models"
)

type Database struct{}

func (d *Database) Connect() *pg.DB {
	databaseCredentials := lib.GetDatabaseCredentials()
	db := pg.Connect(&pg.Options{
		User:     databaseCredentials.PG_USER,
		Password: databaseCredentials.PG_PASSWORD,
		Database: databaseCredentials.PG_DATABASE,
		Addr:     ":5432",
	})

	ctx := context.Background()

	err := db.Ping(ctx)
	if err != nil {
		panic(err)
	}

	return db
}

func (d *Database) InitializeModels(db *pg.DB) {
	err := db.Model(&models.ShortUrls{}).CreateTable(&orm.CreateTableOptions{
		Temp: true,
	})

	if err != nil {
		panic(err)
	}
}
