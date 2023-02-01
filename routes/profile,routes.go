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

func (rc *ProfileRouteController) AuthRoute(rg *gin.RouterGroup) {
	//	router := rg.Group("auth")

}
