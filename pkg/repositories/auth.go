package repositories

import "database/sql"

type Authorization struct {
	db *sql.DB
}

func NewAuthorization(db *sql.DB) *Authorization {
	return &Authorization{
		db: db,
	}
}
