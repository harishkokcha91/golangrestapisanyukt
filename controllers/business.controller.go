package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type BusinessController struct {
	DB *gorm.DB
}

func NewBusinessController(DB *gorm.DB) BusinessController {
	return BusinessController{DB}
}

func (bc *BusinessController) CreateBusiness(ctx *gin.Context) {
	var payload models.CreateBusinessRequest

	// Binding JSON payload to CreateBusinessRequest struct
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assuming "currentUser" is set in middleware and is of type models.User
	// currentUser := ctx.MustGet("currentUser").(models.User)

	// Creating a new business instance
	newBusiness := models.Business{
		ID:           uuid.New(),
		Name:         payload.Name,
		Description:  payload.Description,
		CategoryID:   payload.CategoryID,
		OwnerName:    payload.OwnerName,
		PhoneNumber:  payload.PhoneNumber,
		Email:        payload.Email,
		Website:      payload.Website,
		AddressLine1: payload.AddressLine1,
		AddressLine2: payload.AddressLine2,
		City:         payload.City,
		State:        payload.State,
		PostalCode:   payload.PostalCode,
		Country:      payload.Country,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// Inserting the new business into the database
	result := bc.DB.Create(&newBusiness)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Business with that name or details already exists"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	// Responding with the created business data
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newBusiness})
}

func (bc *BusinessController) UpdateBusiness(ctx *gin.Context) {
	businessID := ctx.Param("businessId")
	// currentUser := ctx.MustGet("currentUser").(models.User)

	var payload models.UpdateBusinessRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var existingBusiness models.Business
	if err := bc.DB.First(&existingBusiness, "id = ?", businessID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No business with that ID exists"})
		return
	}

	if payload.Name != nil {
		existingBusiness.Name = *payload.Name
	}
	if payload.Description != nil {
		existingBusiness.Description = *payload.Description
	}
	if payload.CategoryID != nil {
		existingBusiness.CategoryID = *payload.CategoryID
	}
	if payload.OwnerName != nil {
		existingBusiness.OwnerName = *payload.OwnerName
	}
	if payload.PhoneNumber != nil {
		existingBusiness.PhoneNumber = *payload.PhoneNumber
	}
	if payload.Email != nil {
		existingBusiness.Email = *payload.Email
	}
	if payload.Website != nil {
		existingBusiness.Website = *payload.Website
	}
	if payload.AddressLine1 != nil {
		existingBusiness.AddressLine1 = *payload.AddressLine1
	}
	if payload.AddressLine2 != nil {
		existingBusiness.AddressLine2 = *payload.AddressLine2
	}
	if payload.City != nil {
		existingBusiness.City = *payload.City
	}
	if payload.State != nil {
		existingBusiness.State = *payload.State
	}
	if payload.PostalCode != nil {
		existingBusiness.PostalCode = *payload.PostalCode
	}
	if payload.Country != nil {
		existingBusiness.Country = *payload.Country
	}

	existingBusiness.UpdatedAt = time.Now()

	if err := bc.DB.Save(&existingBusiness).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to update business"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": existingBusiness})
}
func (bc *BusinessController) FindBusinessById(ctx *gin.Context) {
	businessId := ctx.Param("businessId")

	var business models.Business
	result := bc.DB.First(&business, "id = ?", businessId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No business with that ID exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": business})
}

func (bc *BusinessController) GetBusinesses(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		intPage = 1
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		intLimit = 10
	}
	offset := (intPage - 1) * intLimit

	var businesses []models.Business
	results := bc.DB.Limit(intLimit).Offset(offset).Find(&businesses)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(businesses), "data": businesses})
}

func (bc *BusinessController) DeleteBusiness(ctx *gin.Context) {
	businessId := ctx.Param("businessId")

	result := bc.DB.Delete(&models.Business{}, "id = ?", businessId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No business with that ID exists"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
