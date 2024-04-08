package postgresql

import (
	"database/sql"

	"github.com/mystpen/car-catalog-api/internal/model"
)

type PeopleRepository struct {
	db *sql.DB
}

func NewPeopleRepository(db *sql.DB) *PeopleRepository {
	return &PeopleRepository{db: db}
}

func (pr *PeopleRepository) Insert(person *model.Person) error {
	return nil
}
