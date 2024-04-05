package postgresql

import "database/sql"

type PeopleRepository struct {
	db *sql.DB
}

func NewPeopleRepository(db *sql.DB) *PeopleRepository{
	return &PeopleRepository{db: db}
}