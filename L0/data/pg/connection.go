package pg

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var db = connectToDB()

func connectToDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres port=5432  dbname=postgres password=2016001 sslmode=disable")

	if err != nil {
		log.Fatalln(err)

	}
	return db
}
