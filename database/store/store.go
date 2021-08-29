package store

import (
	"github.com/go-pg/pg/v10"
	"log"
)

type User struct {
	Username string `binding:"required,min=5,max=30"`
	Password string `binding:"required,min=7,max=32"`
}

var Users []*User

var db *pg.DB

func SetDBConnection(dbOpts *pg.Options) {
	if dbOpts == nil {
		log.Panicln("DB options can't be nil")
	} else {
		db = pg.Connect(dbOpts)
	}
}

func GetDBConnection() *pg.DB { return db }