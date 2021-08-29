package store

import (
	"github.com/go-pg/pg/v10"
	"os"
)

const (
	addr = "PSQL_ADDR"
	database = "PSQL_DATABASE"
	user = "PSQL_USER"
	password = "PSQL_PASSWORD"
)

func NewDBOptions() *pg.Options {
	return &pg.Options{
		Addr:     os.Getenv(addr),
		Database: os.Getenv(database),
		User:     os.Getenv(user),
		Password: os.Getenv(password),
	}
}