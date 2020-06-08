package util

import (
	"VICCIService/internals/entities"
	"github.com/stretchr/testify/mock"
)

type MockVICCIService struct {
	mock.Mock
}

func (mock *MockVICCIService) GetAllCarLines(tenant string) ([]entities.Carline, error) {
	args := mock.Called(tenant)
	return args.Get(0).([]entities.Carline), args.Error(1)
}

func (mock *MockVICCIService) GetCatalogSummary(tenant string, carline string) (entities.Catalog, error) {
	args := mock.Called(tenant, carline)
	return args.Get(0).(entities.Catalog), args.Error(1)
}

