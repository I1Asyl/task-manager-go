package database

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	// postgres driver
	_ "github.com/lib/pq"
)

// NewConnection creates a new connection from the given configs
func NewConnection() (*sql.DB, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, fmt.Errorf("Failed to parse PORT, err: %v", err)
	}

	// connection string for postgres database
	// host, port, user, password, database name, sslmode, timeout, max connections, max idle connections, max lifetime of connection, max lifetime of idle connection, max lifetime of connection in transaction, max lifetime of idle connection in transaction, max lifetime of connection in prepared statement, max lifetime of idle connection in prepared statement, max lifetime of connection in statement, max lifetime of idle connection in statement, max lifetime of connection in row, max lifetime of idle connection in row, max lifetime of connection in result, max lifetime of idle connection in result, max lifetime of connection in transaction result, max lifetime of idle connection in transaction result, max lifetime of connection in query result, max lifetime of idle connection in query result, max lifetime of connection in query result, max lifetime of idle connection in query result, max lifetime of connection in query result, max lifetime of idle connection in query result, max lifetime of connection in query result, max lifetime of idle connection in
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), port, os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DBNAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("Failed to open connection to Database, err: %v", err)
	}

	return db, nil

}
