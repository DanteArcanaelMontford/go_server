package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	connection := "user=postgres dbname=go_web_test password=1234 host=192.168.200.194 sslmode=disable"
	db, err := sql.Open("postgres", connection)

	if err != nil {
		panic(err.Error())
	}

	return db
}
