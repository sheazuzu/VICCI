package api

import (
	"VICCIService/internals/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type VICCIService interface {
	GetAllCarLines(tenant string) ([]entities.Carline, error)
	GetCatalogSummary(tenant string, carline string) (entities.Catalog, error)
}

type VICCIApi struct {
	VICCIService VICCIService
}

func ProvideVICCIApi(VICCIService VICCIService) *VICCIApi {
	return &VICCIApi{VICCIService: VICCIService}
}


func (VICCIApi *VICCIApi)GetAllCarlines(requestContext *gin.Context){
	tenant := requestContext.Param("vicci_tenant")
	if tenant == ""{
		requestContext.JSON(http.StatusBadRequest, "Return 400, Invalid Request")
		return
	}

	carlines, err :=VICCIApi.VICCIService.GetAllCarLines(tenant)
	if err != nil {
		requestContext.JSON(http.StatusNotFound, "Return 404, No data found")
		return
	}
	requestContext.JSON(http.StatusOK, carlines)
}


func (VICCIApi *VICCIApi)GetCatalogSummary(requestContext *gin.Context){
	tenant := requestContext.Param("vicci_tenant")
	if tenant == ""{
		requestContext.JSON(http.StatusBadRequest, "Return 400, Invalid Request")
		return
	}

	carline := requestContext.Request.URL.Query().Get("carline")
	if carline == ""{
		requestContext.JSON(http.StatusBadRequest, "Return 400, Invalid Request")
		return
	}
	catalog, err := VICCIApi.VICCIService.GetCatalogSummary(tenant ,carline)
	if err != nil {
		requestContext.JSON(http.StatusNotFound, "Return 404, No data found")
		return
	}
	requestContext.JSON(http.StatusOK, catalog)
}



func SetupRouter(router *gin.Engine, VICCIService VICCIService){
	VICCIApi := ProvideVICCIApi(VICCIService)

	router.GET("/tenant/:vicci_tenant/carlines", VICCIApi.GetAllCarlines)
	router.GET("/tenant/:vicci_tenant/catalog", VICCIApi.GetCatalogSummary)

}

