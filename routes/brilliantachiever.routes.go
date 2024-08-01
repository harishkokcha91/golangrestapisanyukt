package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
)

type BrilliantAchieverRouteController struct {
	achieverController controllers.BrilliantAchieverController
}

func NewBrilliantAchieverRouteController(achieverController controllers.BrilliantAchieverController) BrilliantAchieverRouteController {
	return BrilliantAchieverRouteController{achieverController}
}

func (rc *BrilliantAchieverRouteController) BrilliantAchieverRoutes(rg *gin.RouterGroup) {
	router := rg.Group("/achievers")
	router.POST("/", rc.achieverController.CreateAchiever)
	router.GET("/:id", rc.achieverController.GetAchieverByID)
	router.GET("/", rc.achieverController.GetAchievers)
	router.PUT("/:id", rc.achieverController.UpdateAchiever)
	router.DELETE("/:id", rc.achieverController.DeleteAchiever)
}
