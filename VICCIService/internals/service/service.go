package service

import (
	"VICCIService/internals/entities"
)

type VICCIRepository interface {
	FindCarLines(tenant string)([]entities.Carline, error)
	FindCatalogSummary(tenant string, carline string) (entities.Catalog, error)
}

type VICCIService struct {
	VICCIRepository VICCIRepository
}

	func ProvideVICCIService(VICCIRepository VICCIRepository) *VICCIService {
	return &VICCIService{VICCIRepository: VICCIRepository}
}

func (VICCEService *VICCIService) GetAllCarLines(tenant string)  ([]entities.Carline, error) {
	carline, err := VICCEService.VICCIRepository.FindCarLines(tenant)
	if err != nil{
		return nil, err
	}
	return carline, nil

}


func (VICCEService *VICCIService) GetCatalogSummary(tenant string, carline string) (entities.Catalog, error){
	catalog, err := VICCEService.VICCIRepository.FindCatalogSummary(tenant,carline)
	if err != nil{
		return catalog, err
	}
	return catalog, nil
}
