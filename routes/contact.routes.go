// routes/contact.routes.go
package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
	"github.com/wpcodevo/golang-gorm-postgres/middleware"
)

type ContactRouteController struct {
	contactController controllers.ContactController
}

func NewContactRouteController(contactController controllers.ContactController) ContactRouteController {
	return ContactRouteController{contactController}
}

func (cc *ContactRouteController) ContactRoute(rg *gin.RouterGroup) {
	router := rg.Group("contacts")
	router.Use(middleware.DeserializeUser()) // Optional: Use if authentication is required
	router.POST("/", cc.contactController.CreateContact)
	router.GET("/", cc.contactController.FindContacts)
	router.PUT("/:id", cc.contactController.UpdateContact)
	router.GET("/:id", cc.contactController.FindContactById)
	router.DELETE("/:id", cc.contactController.DeleteContact)
}
