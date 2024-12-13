package cmd

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/url_shortner/models"
)

func InitializeModels(db *pg.DB) {
	err := db.Model(&models.ShortUrls{}).CreateTable(&orm.CreateTableOptions{
		Temp: true,
	})

	if err != nil {
		panic(err)
	}
}
