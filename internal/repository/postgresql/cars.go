package postgresql

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/mystpen/car-catalog-api/internal/model"
)

type CarsRepository struct {
	db *sql.DB
}

func NewCarsRepository(db *sql.DB) *CarsRepository {
	return &CarsRepository{db: db}
}

func (cr *CarsRepository) Get(id int64) (*model.CarInfo, error) {
	query := `
	SELECT car_id, regNum, mark, model, year, person_id, name, surname, patronymic
	FROM people INNER JOIN cars
	ON cars.owner_id = people.person_id
	WHERE car_id=$1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var carinfo model.CarInfo

	err := cr.db.QueryRowContext(ctx, query, id).Scan(
		&carinfo.ID,
		&carinfo.RegNum,
		&carinfo.Mark,
		&carinfo.Model,
		&carinfo.Year,
		&carinfo.Owner.ID,
		&carinfo.Owner.Name,
		&carinfo.Owner.Surname,
		&carinfo.Owner.Patronymic,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &carinfo, nil
}

func (cr *CarsRepository) GetAll(filters model.Filters) ([]*model.CarInfo, error) {
	query := `
		SELECT car_id, regNum, mark, model, year, person_id, name, surname, patronymic
		FROM people INNER JOIN cars
		ON cars.owner_id = people.person_id
		WHERE (STRPOS(LOWER(regNum), LOWER($1)) > 0 OR $1 = '')
		AND (LOWER(mark)=LOWER($2)  OR $2 = '')
		AND (LOWER(model)=LOWER($3)  OR $3 = '')
		AND (year=$4  OR $4 = 0)
		ORDER BY car_id ASC
		LIMIT $5 OFFSET $6`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{
		filters.RegNum,
		filters.Mark,
		filters.Model,
		filters.Year,
		filters.PageSize,
		(filters.Page - 1) * filters.PageSize,
	}

	rows, err := cr.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := []*model.CarInfo{}

	for rows.Next() {
		var car model.CarInfo
		err := rows.Scan(
			&car.ID,
			&car.RegNum,
			&car.Mark,
			&car.Model,
			&car.Year,
			&car.Owner.ID,
			&car.Owner.Name,
			&car.Owner.Surname,
			&car.Owner.Patronymic,
		)
		if err != nil {
			return nil, err
		}

		cars = append(cars, &car)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return cars, nil
}

func (cr *CarsRepository) Insert(carInfo *model.CarInfo) error {
	query := `
		INSERT INTO cars (regNum, mark, model, year, owner_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING car_id`

	args := []any{
		carInfo.RegNum,
		carInfo.Mark,
		carInfo.Model,
		carInfo.Year,
		carInfo.Owner.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return cr.db.QueryRowContext(ctx, query, args...).Scan(&carInfo.ID)
}

func (cr *CarsRepository) Update(cars *model.CarInfo) error {
	query := `
		UPDATE cars
		SET regNum = $1, mark = $2, model = $3, year = $4
		WHERE car_id = $5`

	args := []any{
		cars.RegNum,
		cars.Mark,
		cars.Model,
		cars.Year,
		cars.ID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := cr.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (cr *CarsRepository) Delete(id int64) error {
	query := `
	DELETE FROM cars
	WHERE car_id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := cr.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
