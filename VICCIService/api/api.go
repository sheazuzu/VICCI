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

	carlines, err :=VICCIApi.VICCIService.GetAllCarLines(tenant)
	if err != nil {
		requestContext.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	requestContext.JSON(http.StatusOK, carlines)
}


func (VICCIApi *VICCIApi)GetCatalogSummary(requestContext *gin.Context){
	tenant := requestContext.Param("vicci_tenant")
	catalog, err := VICCIApi.VICCIService.GetCatalogSummary(tenant ,requestContext.Request.URL.Query().Get("carline"))
	if err != nil {
		requestContext.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	requestContext.JSON(http.StatusOK, catalog)
}



func SetupRouter(router *gin.Engine, VICCIService VICCIService, Auth entities.Auth){
	VICCIApi := ProvideVICCIApi(VICCIService)

	authGroup:=router.Group("/", gin.BasicAuth(gin.Accounts{
		Auth.User: Auth.Password,
	}))

	authGroup.GET("/tenant/:vicci_tenant/carlines", VICCIApi.GetAllCarlines)
	authGroup.GET("/tenant/:vicci_tenant/catalog", VICCIApi.GetCatalogSummary)

}

