package dataapi

import (
	"log"

	"github.com/go-pg/pg/v10"
)

var db *pg.DB

func setupCockroach() {
	log.Println(config.CockroachURI)
	opt, err := pg.ParseURL(config.CockroachURI)
	if err != nil {
		log.Println(err)
		panic(err)
	}
	db = pg.Connect(opt)
}
