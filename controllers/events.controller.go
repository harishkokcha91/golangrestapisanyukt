package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type EventController struct {
	DB *gorm.DB
}

func NewEventController(DB *gorm.DB) EventController {
	return EventController{DB}
}

func (ec *EventController) CreateEvent(ctx *gin.Context) {
	var payload models.Event
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	now := time.Now()
	payload.CreatedAt = now
	payload.UpdatedAt = now

	result := ec.DB.Create(&payload)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Event with that title already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": payload})
}

func (ec *EventController) UpdateEvent(ctx *gin.Context) {
	eventId := ctx.Param("eventId")

	var payload models.Event
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var existingEvent models.Event
	result := ec.DB.First(&existingEvent, "id = ?", eventId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No event with that ID exists"})
		return
	}

	payload.UpdatedAt = time.Now()
	ec.DB.Model(&existingEvent).Updates(payload)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": existingEvent})
}

func (ec *EventController) FindEventById(ctx *gin.Context) {
	eventId := ctx.Param("eventId")

	var event models.Event
	result := ec.DB.First(&event, "id = ?", eventId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No event with that ID exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": event})
}

func (ec *EventController) FindEvents(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var events []models.Event
	results := ec.DB.Limit(intLimit).Offset(offset).Find(&events)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(events), "data": events})
}

func (ec *EventController) DeleteEvent(ctx *gin.Context) {
	eventId := ctx.Param("eventId")

	result := ec.DB.Delete(&models.Event{}, "id = ?", eventId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No event with that ID exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
