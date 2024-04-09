package service

import (
	"errors"

	"github.com/mystpen/car-catalog-api/internal/model"
	"github.com/mystpen/car-catalog-api/internal/repository/api"
	"github.com/mystpen/car-catalog-api/pkg/logger"
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

	ApiClient interface {
		GetCarInfo(regNum string) (*model.CarInfo, error)
	}
)

type CarCatalogService struct {
	carRepo    CarStorage
	peopleRepo PeopleStorage
	apiClient ApiClient
}

func NewCarCatalogService(carRepo CarStorage, peopleRepo PeopleStorage, apiClient ApiClient) *CarCatalogService {
	return &CarCatalogService{
		carRepo:    carRepo,
		peopleRepo: peopleRepo,
	}
}

func (cs *CarCatalogService) Get(id int64) (*model.CarInfo, error) {
	return cs.carRepo.Get(id)
}

func (cs *CarCatalogService) GetAll(filters model.Filters) ([]*model.CarInfo, error) {
	return cs.carRepo.GetAll(filters)
}

func (cs *CarCatalogService) InsertRegNums(regNums []string, cars *[]model.CarInfo) error {
	for _, regNum := range regNums {
		carInfo, err := cs.apiClient.GetCarInfo(regNum)
		if err != nil{
			if errors.Is(err, api.ErrBadRequest) {
				logger.PrintDebug("not added regNum", map[string]any{
					"regNum": regNum,
					"error": err,
				})
			} else{
				return err
			}
		}

		if carInfo != nil {
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

	}
	return nil
}

func (cs *CarCatalogService) Update(cars *model.CarInfo) error {
	return cs.carRepo.Update(cars)
}

func (cs *CarCatalogService) Delete(id int64) error {
	return cs.carRepo.Delete(id)
}
