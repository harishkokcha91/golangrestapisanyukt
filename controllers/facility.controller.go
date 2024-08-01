package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type FacilityController struct {
	DB *gorm.DB
}

func NewFacilityController(DB *gorm.DB) FacilityController {
	return FacilityController{DB}
}

func (fc *FacilityController) CreateFacility(ctx *gin.Context) {
	var payload models.Facility

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	result := fc.DB.Create(&payload)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": payload})
}

func (fc *FacilityController) UpdateFacility(ctx *gin.Context) {
	facilityID := ctx.Param("id")
	var payload models.Facility

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var facility models.Facility
	result := fc.DB.First(&facility, "id = ?", facilityID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Facility not found"})
		return
	}

	fc.DB.Model(&facility).Updates(payload)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": facility})
}

func (fc *FacilityController) GetFacilityByID(ctx *gin.Context) {
	facilityID := ctx.Param("id")

	var facility models.Facility
	result := fc.DB.First(&facility, "id = ?", facilityID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Facility not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": facility})
}

func (fc *FacilityController) ListFacilities(ctx *gin.Context) {
	var facilities []models.Facility
	result := fc.DB.Find(&facilities)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": facilities})
}

func (fc *FacilityController) DeleteFacility(ctx *gin.Context) {
	facilityID := ctx.Param("id")

	result := fc.DB.Delete(&models.Facility{}, "id = ?", facilityID)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Facility not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
