package delivery

import (
	"time"

	"github.com/mystpen/car-catalog-api/internal/model"
	"github.com/mystpen/car-catalog-api/pkg/validator"
)

func ValidateFilters(v *validator.Validator, f model.Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must be a maximum of 10 million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")
}

func ValidateRegNums(v *validator.Validator, regNums []string) {
	for _, regNum := range regNums {
		if !v.Matches(regNum, validator.RegNumRX) {
			v.AddError("regNum", "invalid format for regNum")
			break
		}
	}
}

func ValidateCarInfo(v *validator.Validator, car *model.CarInfo) {
	v.Check(v.Matches(car.RegNum, validator.RegNumRX), "regNum", "invalid format for regNum")

	v.Check(car.Mark != "", "mark", "must be provided")
	v.Check(len(car.Mark) <= 500, "mark", "must not be more than 500 bytes long")

	v.Check(car.Model != "", "model", "must be provided")
	v.Check(len(car.Model) <= 500, "model", "must not be more than 500 bytes long")

	if car.Year != 0 {
		v.Check(car.Year >= 1000, "year", "must be greater than 1000")
		v.Check(car.Year <= int(time.Now().Year()), "year", "must not be in the future")
	}

	v.Check(car.Owner.Name != "", "owner_name", "must be provided")
	v.Check(len(car.Owner.Name) <= 500, "owner_name", "must not be more than 500 bytes long")

	v.Check(car.Owner.Surname != "", "owner_surname", "must be provided")
	v.Check(len(car.Owner.Surname) <= 500, "owner_surname", "must not be more than 500 bytes long")

	if car.Owner.Patronymic != ""{
		v.Check(len(car.Owner.Patronymic) <= 500, "owner_patronymic", "must not be more than 500 bytes long")
	}
}
