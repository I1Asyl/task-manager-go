package database

import (
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
)

// NewConnection creates a new connection from the given configs
func NewConnection(dsn string) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to open connection to Database, err: %v", err)
	}

	return db, nil

}
