package service

import (
	"github.com/mystpen/car-catalog-api/internal/model"
)

type (
	CarStorage interface {
		Get(id int64) (*model.CarInfo, error)
		GetAll(filters model.Filters) ([]*model.CarInfo, error)
		Insert(*model.CarInfo) error
		Update(cars *model.CarInfo) error
		Delete(id int64) error
	}

	PeopleStorage interface {
		Insert(*model.Person) error
	}
)

type CarCatalogService struct {
	carRepo    CarStorage
	peopleRepo PeopleStorage
}

func NewCarCatalogService(carRepo CarStorage, peopleRepo PeopleStorage) *CarCatalogService {
	return &CarCatalogService{
		carRepo:    carRepo,
		peopleRepo: peopleRepo,
	}
}

func (cs *CarCatalogService) Get(id int64) (*model.CarInfo, error){
	return cs.carRepo.Get(id)
}

func (cs *CarCatalogService) GetAll(filters model.Filters) ([]*model.CarInfo, error) {
	return cs.carRepo.GetAll(filters)
}

func (cs *CarCatalogService) InsertRegNums(regNums []string, cars *[]model.CarInfo) error {
	for range regNums {
		var carInfo *model.CarInfo
		// get info from API

		err := cs.peopleRepo.Insert(&carInfo.Owner)
		if err != nil {
			return err
		}

		err = cs.carRepo.Insert(carInfo)
		if err != nil {
			return err
		}

		*cars = append(*cars, *carInfo)
	}
	return nil
}

func (cs *CarCatalogService) Update(cars *model.CarInfo) error{
	return cs.carRepo.Update(cars)
}

func (cs *CarCatalogService) Delete(id int64) error{
	return cs.carRepo.Delete(id)
}