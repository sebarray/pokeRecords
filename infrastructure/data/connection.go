package data

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnection() *sqlx.DB {

	db, err := sqlx.Connect("postgres", "user=myuser dbname=POKERECORD password=mypassword sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	return db

}
