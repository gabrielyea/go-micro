package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

var db *sql.DB

// ConnectDB Creates a connection to a PostgreSQL database using the global constants.
func ConnectDB() (*sql.DB, error) {
	var err error

	db, err = NewPgConn()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}

func NewPgConn() (*sql.DB, error) {
	dsn := os.Getenv("DNS")

	var err error
	db, err = sql.Open("pgx", dsn)

	if err != nil {
		return nil, fmt.Errorf("couldn't prepare connection to database %s, %s", "db", err)
	}

	return db, nil
}
