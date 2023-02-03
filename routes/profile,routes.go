package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
)

type ProfileRouteController struct {
	profileController controllers.ProfileController
}

func NewProfileRouteController(profileController controllers.ProfileController) ProfileRouteController {
	return ProfileRouteController{profileController}
}

func (rc *ProfileRouteController) ProfileRoute(rg *gin.RouterGroup) {
	router := rg.Group("profile")
	router.POST("/create", rc.profileController.CreateProfile)
	router.POST("/getProfiles", rc.profileController.GetProfiles)
	router.POST("/getProfile", rc.profileController.GetProfile)
	router.POST("/getMyProfiles", rc.profileController.GetMyProfiles)
	router.POST("/createFavProfiles", rc.profileController.CreateFavProfiles)
	router.POST("/getFavProfiles", rc.profileController.GetFavProfiles)
}
