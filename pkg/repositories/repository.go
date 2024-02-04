package repositories

import "database/sql"

type auth interface {
}

type static interface {
	ReturnStatus() string
}

// Repository structure to c
type Repository struct {
	auth
	static
}

// New returns a new repository with relevant methods configured
func New(db *sql.DB) *Repository {
	return &Repository{
		auth:   NewAuthorization(db),
		static: NewStatic(db),
	}
}

// use sqlc for generating db access code in Go
