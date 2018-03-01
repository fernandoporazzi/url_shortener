package database

import (
	"database/sql"
)

// Connect create a new database connection
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/url_shortener?charset=utf8")
	if err != nil {
		panic(err)
	}

	return db
}