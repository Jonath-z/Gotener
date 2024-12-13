package db

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/url_shortner/lib"
)

type Database struct {
	Db *pg.DB
}

func (d *Database) Connect() {
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
	d.Db = db
}
