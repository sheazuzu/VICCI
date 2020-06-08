package entities

type Catalog struct {
	Name string `json:"name"`
	Code string`json:"code"`
	Salesgroup []Salesgroup
}

type Salesgroup struct {
	Name string `json:"name"`
	Code string`json:"code"`
	Models []Model
}


type Model struct {
	Name string `json:"name"`
	Code string`json:"code"`
	Version string `json:"version"`
	ModelYear string `json:"year"`
}