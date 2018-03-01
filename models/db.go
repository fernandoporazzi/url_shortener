package models

import "database/sql"

var db *sql.DB

// Init database
func Init(database *sql.DB) {
	db = database
}