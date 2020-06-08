package service

import (
	"VICCIService/internals/entities"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockVICCIRepository struct {
	 mock.Mock
}

func (mock *MockVICCIRepository) FindCarLines (tenant string) ([]entities.Carline, error){
	args := mock.Called(tenant)
	return args.Get(0).([]entities.Carline), args.Error(1)
}

func (mock *MockVICCIRepository) FindCatalogSummary(tenant string, carline string) (entities.Catalog, error) {
	args := mock.Called(tenant, carline)
	return args.Get(0).(entities.Catalog), args.Error(1)
}

func TestVICCIService_GetAllCarLines(t *testing.T) {
	const findCarLines = "FindCarLines"
	const tenant = "okapi-audi-de-de"

	carlines := []entities.Carline{
		{
			Name: "AUDI RS 7",
			Code: "11111",
		},
	}

	t.Run("Get All Carlines ", func(t *testing.T){
		mockCarline := MockVICCIRepository{}
		mockCarline.On(findCarLines,tenant).Return(carlines, nil)
		service :=ProvideVICCIService(&mockCarline)
		result, err := service.GetAllCarLines(tenant)
		assert.NoError(t,err)
		assert.Equal(t, carlines, result)
		mockCarline.AssertNumberOfCalls(t,findCarLines,1)
	})

	t.Run("Get All Carlines with return Erroe ", func(t *testing.T){
		mockCarline := MockVICCIRepository{}
		mockCarline.On(findCarLines,tenant).Return(carlines, errors.New("test"))
		service :=ProvideVICCIService(&mockCarline)
		_, err := service.GetAllCarLines(tenant)
		assert.Error(t,err)
		mockCarline.AssertNumberOfCalls(t,findCarLines,1)
	})

}


func TestVICCIService_GetCatalogSummary(t *testing.T) {
	const findCatalogSummary = "FindCatalogSummary"
	const tenant = "okapi-audi-de-de"
	const carline = "30010"

	model := []entities.Model{
		{
			Name:      "AUDI RS 7",
			Code:      "33333",
			Version:   "SportWagon",
			ModelYear: "2020",
		},
	}
	salesgroup := []entities.Salesgroup{
		{
			Name:   "AUDI RS 7",
			Code:   "11112",
			Models: model,
		},
	}
	catalog := entities.Catalog{

		Name:       "AUDI RS 7",
		Code:       "11111",
		Salesgroup: salesgroup,
	}

	t.Run("Get All Catalog Overview", func(t *testing.T) {
		mockCatalog := MockVICCIRepository{}
		mockCatalog.On(findCatalogSummary, tenant, carline).Return(catalog, nil)
		service := ProvideVICCIService(&mockCatalog)
		result, err := service.GetCatalogSummary(tenant,carline)
		assert.NoError(t, err)
		assert.Equal(t, catalog, result)
		mockCatalog.AssertNumberOfCalls(t,findCatalogSummary,1)
	})

	t.Run("Get All Catalog Overview with return Error", func(t *testing.T) {
		mockCatalog := MockVICCIRepository{}
		mockCatalog.On(findCatalogSummary, tenant, carline).Return(catalog, errors.New("test"))
		service := ProvideVICCIService(&mockCatalog)
		_, err := service.GetCatalogSummary(tenant,carline)
		assert.Error(t, err)
		mockCatalog.AssertNumberOfCalls(t,findCatalogSummary,1)
	})

}