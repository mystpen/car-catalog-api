package service

type (
	CarStorage interface {
	}

	PeopleStorage interface {
	}
)

type CarCatalogService struct {
	carRepo CarStorage
	peopleRepo PeopleStorage
}

func NewCarCatalogService(carRepo CarStorage, peopleRepo PeopleStorage) *CarCatalogService{
	return &CarCatalogService{
		carRepo: carRepo,
		peopleRepo: peopleRepo,
	}
}