package database

import (
	"database/sql"
	"fmt"
)

const (
	DB_DRIVER = "postgres"
	DB_CONNECTION = "user=postgres dbname=postgres password=postgres sslmode=disable"
)

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open(DB_DRIVER, DB_CONNECTION)
	fmt.Println(err)
	db.Ping()
	return db, err;
}
