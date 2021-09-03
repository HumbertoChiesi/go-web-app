package db

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Connect opens the conection with the DB and returns it
func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.StringConnectionDB)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
