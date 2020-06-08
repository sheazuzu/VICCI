package repository

import (
	"VICCIService/internals/entities"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type VICCIRepository struct {
	url string
}

func ProvideVICCIRepository(url string) *VICCIRepository {

	return &VICCIRepository{url}
}


func (VICCIRepository *VICCIRepository)FindCarLines(tenant string) ([]entities.Carline, error){
	var Carlines []entities.Carline

	url := VICCIRepository.url +"catalogue/carlines?tenant=" + tenant
	response, _ :=http.Get(url)
	head, _ := http.Head(url)
	body, _ := ioutil.ReadAll(response.Body)

	if head.Status == "400 " {
		return nil, errors.New("Status 400, Bad Request")
	}
	if head.Status == "404 " {
		return nil, errors.New("Status 404, No data found")
	}

	defer response.Body.Close()
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

	return  Carlines, nil
}

func (VICCIRepository *VICCIRepository)FindCatalogSummary(tenant string, carline string) (entities.Catalog, error) {
	var catalog entities.Catalog
	var salesGroups []entities.Salesgroup
	var Models []entities.Model

	salesgroupsUrl := VICCIRepository.url + "catalogue/salesgroups?tenant=" + tenant +
		 "&carlineKey=" + carline + "&fetchPrices=false&fetchMedia=false&addErrorMps=false"

	response, _ :=http.Get(salesgroupsUrl)
	head, _ := http.Head(salesgroupsUrl)

	salesgroupsBody, _ := ioutil.ReadAll(response.Body)

	if head.Status == "400 " {
		return catalog, errors.New("Status 400, Bad Request")
	}
	if head.Status == "404 " {
		return catalog, errors.New("Status 404, No data found")
	}
	if head.Status == "204 " {
		return catalog, errors.New("Status 204, No Content")
	}

	defer response.Body.Close()
	var unmarshalSalesGroupsBody map[string]interface{}
	json.Unmarshal([]byte(salesgroupsBody), &unmarshalSalesGroupsBody)


	SalesGroupsInterfaceArray	:=	unmarshalSalesGroupsBody["salesgroups"].([]interface{})
	for _, u := range SalesGroupsInterfaceArray{
		SalesGroupsMap := u.(map[string]interface{})
		var salesGroup entities.Salesgroup
		salesGroup.Code = SalesGroupsMap["code"].(string)
		salesGroup.Name = SalesGroupsMap["name"].(string)

		modelsUrl := VICCIRepository.url + "catalogue/models?tenant=" + tenant + "&salesgroupKey=" + salesGroup.Code +
			"&fetchPrices=false&fetchMandatory=false&fetchTechnical=false&fetchEco=false&fetchWltp=false&fetchMedia=false&addErrorMps=false"

		response, _ :=http.Get(modelsUrl)
		head, _ := http.Head(modelsUrl)
		if head.Status == "400 " {
			return catalog, errors.New("Status 400, Bad Request")
		}
		if head.Status == "404 " {
			return catalog, errors.New("Status 404, No data found")
		}
		if head.Status == "204 " {
			return catalog, errors.New("Status 204, No Content")
		}


		modelsBody, _ := ioutil.ReadAll(response.Body)

		var unmarshalModelsBody map[string]interface{}
		json.Unmarshal([]byte(modelsBody), &unmarshalModelsBody)

		ModelsInterfaceArray := unmarshalModelsBody["models"].([]interface{})
		 for _, v := range ModelsInterfaceArray{
		 	ModelsMap := v.(map[string]interface{})
			 var Model entities.Model
		 	Model.Name = ModelsMap["name"].(string)
		 	Model.Code = ModelsMap["code"].(string)
		 	Model.Version = ModelsMap["version"].(string)
		 	Model.ModelYear = ModelsMap["year"].(string)
		 	catalog.Name = ModelsMap["carlineName"].(string)
		 	Models = append( Models, Model)
		 }

		salesGroup.Models = Models
		salesGroups = append(salesGroups, salesGroup)

	}
	catalog.Code = carline
	catalog.Salesgroup = salesGroups

	return  catalog, nil
}