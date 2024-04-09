package postgresql

import (
	"context"
	"database/sql"
	"time"

	"github.com/mystpen/car-catalog-api/internal/model"
)

type PeopleRepository struct {
	db *sql.DB
}

func NewPeopleRepository(db *sql.DB) *PeopleRepository {
	return &PeopleRepository{db: db}
}

func (pr *PeopleRepository) Insert(person *model.Person) error {
	query := `
		INSERT INTO people (name, surname, patronymic)
		VALUES ($1, $2, $3)
		RETURNING person_id`

	args := []any{person.Name, person.Surname, person.Patronymic}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return pr.db.QueryRowContext(ctx, query, args...).Scan(&person.ID)
}
