package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
)

type FacilityRouteController struct {
	facilityController controllers.FacilityController
}

func NewFacilityRouteController(facilityController controllers.FacilityController) FacilityRouteController {
	return FacilityRouteController{facilityController}
}

func (rc *FacilityRouteController) FacilityRoute(rg *gin.RouterGroup) {
	router := rg.Group("facilities")

	router.POST("/", rc.facilityController.CreateFacility)
	router.PUT("/:id", rc.facilityController.UpdateFacility)
	router.GET("/:id", rc.facilityController.GetFacilityByID)
	router.GET("/", rc.facilityController.ListFacilities)
	router.DELETE("/:id", rc.facilityController.DeleteFacility)
}
