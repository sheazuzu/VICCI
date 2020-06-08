package repository

import (
	"VICCIService/internals/entities"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

type VICCIRepository struct {
	url string
}

func ProvideVICCIRepository(url string) *VICCIRepository {

	return &VICCIRepository{url}
}


func (VICCIRepository *VICCIRepository)FindCarLines(tenant string) ([]entities.Carline, error){
	var Carlines []entities.Carline
	username ,_ := os.LookupEnv("AUTH_USER")
	password, _ := os.LookupEnv("AUTH_PASSWORD")

	url := VICCIRepository.url +"catalogue/carlines?tenant=" + tenant

	client := &http.Client{}
	req, err := http.NewRequest("GET",url,nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(username,password)
	response, err :=client.Do(req)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200{
		return nil, errors.New("false input")
	}


	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}


	defer response.Body.Close()
	var unmarshalBody map[string]interface{}
	json.Unmarshal([]byte(body), &unmarshalBody)

	CarlineInterfaceArray	:=	unmarshalBody["carlines"].([]interface{})
	if CarlineInterfaceArray == nil {
		return nil, errors.New("false carline")
	}

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

	username ,_ := os.LookupEnv("AUTH_USER")
	password, _ := os.LookupEnv("AUTH_PASSWORD")

	salesgroupsUrl := VICCIRepository.url + "catalogue/salesgroups?tenant=" + tenant +
		 "&carlineKey=" + carline + "&fetchPrices=false&fetchMedia=false&addErrorMps=false"

	client := &http.Client{}
	req, err := http.NewRequest("GET",salesgroupsUrl,nil)
	if err != nil {
		panic(err)
	}

	req.SetBasicAuth(username,password)
	response, err :=client.Do(req)
	if err != nil {
		panic(err)
	}

	if response.StatusCode != 200{
		return entities.Catalog{}, errors.New("false input")
	}

	salesgroupsBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	/*
	if head.Status == "400 " {
		return catalog, errors.New("Status 400, Bad Request")
	}
	if head.Status == "404 " {
		return catalog, errors.New("Status 404, No data found")
	}
	if head.Status == "204 " {
		return catalog, errors.New("Status 204, No Content")
	}
	*/


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

		req, err := http.NewRequest("GET",modelsUrl,nil)
		if err != nil {
			panic(err)
		}

		req.SetBasicAuth(username,password)
		response, err :=client.Do(req)
		if err != nil {
			panic(err)
		}

		if response.StatusCode != 200{
			return entities.Catalog{}, errors.New("false input")
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

