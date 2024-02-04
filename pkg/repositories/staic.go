package repositories

import "database/sql"

type Static struct {
	db *sql.DB
}

func NewStatic(db *sql.DB) *Static {
	return &Static{
		db: db,
	}
}

// DoSomething function is a method of Repository
func (s *Static) ReturnStatus() string {
	return "pong"
}
