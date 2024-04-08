package postgresql

import (
	"database/sql"

	"github.com/mystpen/car-catalog-api/internal/model"
)

type CarsRepository struct {
	db *sql.DB
}

func NewCarsRepository(db *sql.DB) *CarsRepository {
	return &CarsRepository{db: db}
}

func (cr *CarsRepository) Get(id int64) (*model.CarInfo, error){
	return nil, nil
}

func (cr *CarsRepository) GetAll(filters model.Filters) ([]*model.CarInfo, error){
	return nil, nil
}

func (cr *CarsRepository) Insert(carInfo *model.CarInfo) error{
	return nil
}

func (cr *CarsRepository) Update(cars *model.CarInfo) error{
	return nil
}

func (cr *CarsRepository) Delete(id int64) error{
	return nil
}