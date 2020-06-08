package api

import (
	"VICCIService/internals/entities"
	"VICCIService/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVICCIApi_GetAllCarlines(t *testing.T) {
	const getAllCarLines = "GetAllCarLines"
	const tenant = "okapi-audi-de-de"

	carlines := []entities.Carline{
		{
		Name: "AUDI RS 7",
		Code: "11111",
		},
	}

	t.Run("Get All Carlines with tenant", func(t *testing.T) {
		mockCarline := util.MockVICCIService{}
		mockCarline.On(getAllCarLines,tenant).Return(carlines,nil)
		router := gin.Default()
		SetupRouter(router, &mockCarline)
		result := util.PerformRequest(router, "GET","/tenant/okapi-audi-de-de/carlines",nil)
		expected := `[{"name":"AUDI RS 7","code":"11111"}]`
		assert.Equal(t, http.StatusOK, result.Code)
		assert.Equal(t, expected, result.Body.String())
		mockCarline.AssertNumberOfCalls(t, getAllCarLines, 1)
	})
	t.Run("Get All Carlines return Error", func(t *testing.T) {
		mockCarline := util.MockVICCIService{}
		mockCarline.On(getAllCarLines,tenant).Return(carlines, errors.New("test"))
		router := gin.Default()
		SetupRouter(router, &mockCarline)
		result := util.PerformRequest(router, "GET","/tenant/okapi-audi-de-de/carlines",nil)
		assert.Equal(t, http.StatusInternalServerError, result.Code)
		mockCarline.AssertNumberOfCalls(t, getAllCarLines, 1)

	})
}

func TestVICCIApi_GetCatalogSummary(t *testing.T) {
	const getCatalogSummary = "GetCatalogSummary"
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

		Name: "AUDI RS 7",
		Code: "11111",
		Salesgroup: salesgroup,
	}

	t.Run("Get All Catalogue Overview", func(t *testing.T){
		mockCatalog := util.MockVICCIService{}
		mockCatalog.On(getCatalogSummary,tenant,carline).Return(catalog,nil)
		router := gin.Default()
		SetupRouter(router, &mockCatalog)
		result := util.PerformRequest(router, "GET","/tenant/okapi-audi-de-de/catalog?carline=30010",nil)
		expected := `{"name":"AUDI RS 7","code":"11111","Salesgroup":[{"name":"AUDI RS 7","code":"11112","Models":[{"name":"AUDI RS 7","code":"33333","version":"SportWagon","year":"2020"}]}]}`
		assert.Equal(t, http.StatusOK, result.Code)
		assert.Equal(t, expected, result.Body.String())
		mockCatalog.AssertNumberOfCalls(t, getCatalogSummary, 1)

	})

	t.Run("Get All Catalogue Overview with return Error", func(t *testing.T){
		mockCatalog := util.MockVICCIService{}
		mockCatalog.On(getCatalogSummary,tenant,carline).Return(catalog,errors.New("test"))
		router := gin.Default()
		SetupRouter(router, &mockCatalog)

		result := util.PerformRequest(router, "GET","/tenant/okapi-audi-de-de/catalog?carline=30010",nil)

		assert.Equal(t, http.StatusInternalServerError, result.Code)
		mockCatalog.AssertNumberOfCalls(t, getCatalogSummary, 1)

	})


}