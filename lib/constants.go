package lib

import "os"

type DatabaseCredentials struct {
	PG_USER     string
	PG_HOST     string
	PG_PASSWORD string
	PG_DATABASE string
}

func GetDatabaseCredentials() DatabaseCredentials {
	return DatabaseCredentials{
		PG_USER:     os.Getenv("PG_USER"),
		PG_HOST:     os.Getenv("PG_HOST"),
		PG_PASSWORD: os.Getenv("PG_PASSWORD"),
		PG_DATABASE: os.Getenv("PG_DATABASE"),
	}
}
