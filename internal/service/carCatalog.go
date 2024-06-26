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
		Update(*model.Person) error
	}

	ApiClient interface {
		GetCarInfo(regNum string) (*model.CarInfo, error)
	}
)

type CarCatalogService struct {
	carRepo    CarStorage
	peopleRepo PeopleStorage
	apiClient  ApiClient
}

func NewCarCatalogService(carRepo CarStorage, peopleRepo PeopleStorage, apiClient ApiClient) *CarCatalogService {
	return &CarCatalogService{
		carRepo:    carRepo,
		peopleRepo: peopleRepo,
		apiClient:  apiClient,
	}
}

func (cs *CarCatalogService) Get(id int64) (*model.CarInfo, error) {
	return cs.carRepo.Get(id)
}

func (cs *CarCatalogService) GetAll(filters model.Filters) ([]*model.CarInfo, error) {
	return cs.carRepo.GetAll(filters)
}

func (cs *CarCatalogService) InsertRegNums(regNums []string) ([]*model.CarInfo, error) {
	var cars []*model.CarInfo
	for _, regNum := range regNums {

		carInfo, err := cs.apiClient.GetCarInfo(regNum)

		logger.PrintDebug("info from external API", map[string]any{
			"carInfo": carInfo,
		})

		if err != nil {
			if errors.Is(err, api.ErrBadRequest) {
				logger.PrintDebug("not added regNum", map[string]any{
					"regNum": regNum,
					"error":  err,
				})
			} else {
				return nil, err
			}
		}

		if carInfo != nil {
			err := cs.peopleRepo.Insert(&carInfo.Owner)
			if err != nil {
				return nil, err
			}

			err = cs.carRepo.Insert(carInfo)
			if err != nil {
				return nil, err
			}

			cars = append(cars, carInfo)
		}

	}
	return cars, nil
}

func (cs *CarCatalogService) Update(cars *model.CarInfo) error {
	if cars.Owner != (model.Person{}) {
		err := cs.peopleRepo.Update(&cars.Owner)
		if err != nil {
			return err
		}
	}
	return cs.carRepo.Update(cars)
}

func (cs *CarCatalogService) Delete(id int64) error {
	return cs.carRepo.Delete(id)
}
