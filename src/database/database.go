package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Connect open a connection to the database and return it, or an error if so
func Connect() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", config.DBConnectionString)
	if err != nil {
		return
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return
	}

	return
}
