package postgresql

import "database/sql"

type CarsRepository struct {
	db *sql.DB
}

func NewCarsRepository(db *sql.DB) *CarsRepository{
	return &CarsRepository{db: db}
}