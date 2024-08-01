package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type BrilliantAchieverController struct {
	DB *gorm.DB
}

func NewBrilliantAchieverController(DB *gorm.DB) BrilliantAchieverController {
	return BrilliantAchieverController{DB}
}

func (bac *BrilliantAchieverController) CreateAchiever(ctx *gin.Context) {
	var payload models.BrilliantAchiever

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	payload.CreatedAt = time.Now()
	if result := bac.DB.Create(&payload); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": payload})
}

func (bac *BrilliantAchieverController) GetAchieverByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var achiever models.BrilliantAchiever
	if result := bac.DB.First(&achiever, id); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Achiever not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": achiever})
}

func (bac *BrilliantAchieverController) GetAchievers(ctx *gin.Context) {
	var achievers []models.BrilliantAchiever

	if result := bac.DB.Find(&achievers); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": achievers})
}

func (bac *BrilliantAchieverController) UpdateAchiever(ctx *gin.Context) {
	id := ctx.Param("id")

	var payload models.BrilliantAchiever
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var achiever models.BrilliantAchiever
	if result := bac.DB.First(&achiever, id); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Achiever not found"})
		return
	}

	if result := bac.DB.Model(&achiever).Updates(payload); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": achiever})
}

func (bac *BrilliantAchieverController) DeleteAchiever(ctx *gin.Context) {
	id := ctx.Param("id")

	if result := bac.DB.Delete(&models.BrilliantAchiever{}, id); result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Achiever not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
