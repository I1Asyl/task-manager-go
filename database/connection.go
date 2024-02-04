package database

import (
	"database/sql"
	"fmt"

	// postgres driver
	"github.com/I1Asyl/task-manager-go/configuration"
	_ "github.com/lib/pq"
)

// NewConnection creates a new connection from the given configs
func NewConnection(dbConfig configuration.DatabaseConfiguration) (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Name)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("Failed to open connection to Database, err: %v", err)
	}

	return db, nil

}
