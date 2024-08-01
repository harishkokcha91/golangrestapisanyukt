package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
)

type EventRouteController struct {
	eventController controllers.EventController
}

func NewEventRouteController(eventController controllers.EventController) EventRouteController {
	return EventRouteController{eventController}
}

func (erc *EventRouteController) EventRoute(rg *gin.RouterGroup) {
	router := rg.Group("events")

	router.POST("/", erc.eventController.CreateEvent)
	router.PUT("/:eventId", erc.eventController.UpdateEvent)
	router.GET("/:eventId", erc.eventController.FindEventById)
	router.GET("/", erc.eventController.FindEvents)
	router.DELETE("/:eventId", erc.eventController.DeleteEvent)
}
