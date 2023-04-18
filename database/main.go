package database

import (
	_ "github.com/lib/pq"
	"github.com/pocketbase/dbx"
)

var Db *dbx.DB

func Ping() {
	db, err := dbx.Open("postgres", "host=localhost port=5433 user=bon password=bon dbname=lenslocked sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db = db
}
