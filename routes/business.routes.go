package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
)

type BusinessRouteController struct {
	businessController controllers.BusinessController
}

func NewBusinessRouteController(businessController controllers.BusinessController) BusinessRouteController {
	return BusinessRouteController{businessController}
}

func (bc *BusinessRouteController) BusinessRoute(rg *gin.RouterGroup) {
	router := rg.Group("businesses")
	// router.Use(middleware.DeserializeUser())
	router.POST("/", bc.businessController.CreateBusiness)
	router.GET("/", bc.businessController.GetBusinesses)
	router.GET("/:businessId", bc.businessController.FindBusinessById)
	router.PUT("/:businessId", bc.businessController.UpdateBusiness)
	router.DELETE("/:businessId", bc.businessController.DeleteBusiness)
}
