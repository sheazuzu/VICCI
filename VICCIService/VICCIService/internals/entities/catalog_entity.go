package entities

type Catalog struct {
	Name string `json:"name"`
	Code string`json:"code"`
	Salesgroup []Salesgroup `json:"salesgroups"`
}

type Salesgroup struct {
	Name string `json:"name"`
	Code string`json:"code"`
	Models []Model `json:"models"`
}


type Model struct {
	Name string `json:"name"`
	Code string`json:"code"`
	Version string `json:"version"`
	ModelYear string `json:"year"`
}