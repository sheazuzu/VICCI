package repository

import (
	"VICCIService/internals/entities"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestVICCIRepository_FindCarLines(t *testing.T) {
	const tenant = "okapi-seat-es-es"
	const fakeTenant = "okapi-seat-es-de"
	const url = "https://ingress.venividivicci.de/vis/"

	execpted := []entities.Carline{{"Ibiza","80320"},{"Leon","80335"},{"Arona","80370"},{"Ateca","80380"},
	 	{"Alhambra","80390"},{"Tarraco","80400"},{"Mii electric","80440"},{"Nuevo Leon","80460"},{"Cupra","81000"},
	}
	//var Carlines []entities.Carline
	repo := ProvideVICCIRepository(url)
	t.Run("get All Carlines by calling the url", func(t *testing.T) {
/*
		url := "https://ingress.venividivicci.de/vis/catalogue/carlines?tenant=" + tenant
		response, _ :=http.Get(url)
		head, _ := http.Head(url)
		body, _ := ioutil.ReadAll(response.Body)
		var unmarshalBody map[string]interface{}
		json.Unmarshal([]byte(body), &unmarshalBody)

		CarlineInterfaceArray	:=	unmarshalBody["carlines"].([]interface{})
		for _, u := range CarlineInterfaceArray{
			CarlineMap := u.(map[string]interface{})
			var carline entities.Carline
			carline.Code = CarlineMap["code"].(string)
			carline.Name = CarlineMap["name"].(string)
			Carlines = append( Carlines, carline)
		}
*/

		got, _ := repo.FindCarLines(tenant)
		assert.Equal(t,execpted,got )
		//assert.Equal(t,"200 ",head.Status)
	})

	t.Run("get All Carlines by calling the url with return Error", func(t *testing.T) {

		url := url + "catalogue/carlines?tenant=" + fakeTenant
		head, _ := http.Head(url)

		_, err := repo.FindCarLines(fakeTenant)
		assert.Error(t,err)
		assert.Equal(t,"400 ",head.Status )
	})


}

func TestVICCIRepository_FindCatalogSummary(t *testing.T) {
	const tenant = "okapi-audi-de-de"
	const carline = "35639"
	const fakeTenant = "okapi-audi-es-de"
	const url = "https://ingress.venividivicci.de/vis/"

	repo := ProvideVICCIRepository(url)
	model :=[]entities.Model{{"RS 7 Sportback   441(600) kW(PS) tiptronic","4KARCA","1","2020"}}
	salesgroup :=[]entities.Salesgroup{{"Audi RS 7 Sportback","50980",model}}

	t.Run("get all catalogs by calling the url", func(t *testing.T) {


		execpted:= entities.Catalog{"RS 7 Sportback","35639",salesgroup}
		got, _ := repo.FindCatalogSummary(tenant,carline)
		assert.Equal(t,execpted,got)
	})

	t.Run("get all catalogs by calling the url with return Error 400", func(t *testing.T) {

		testUrl := url + "catalogue/models?tenant=" + fakeTenant + "&salesgroupKey=" + carline +
			"&fetchPrices=false&fetchMandatory=false&fetchTechnical=false&fetchEco=false&fetchWltp=false&fetchMedia=false&addErrorMps=false"

		head, _ := http.Head(testUrl)
		_, err := repo.FindCatalogSummary(fakeTenant,carline)
		assert.Error(t,err)
		assert.Equal(t,"400 ",head.Status )
	})




}