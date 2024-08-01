// controllers/contact.controller.go
package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type ContactController struct {
	DB *gorm.DB
}

func NewContactController(DB *gorm.DB) ContactController {
	return ContactController{DB}
}

func (cc *ContactController) CreateContact(ctx *gin.Context) {
	// currentUser, _ := ctx.Get("currentUser") // Get current user from context
	var contact models.Contact

	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	// contact.CreatedBy = currentUser.(uint) // Assuming currentUser ID is of type uint
	contact.CreatedBy = 1
	result := cc.DB.Create(&contact)
	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": contact})
}

func (cc *ContactController) UpdateContact(ctx *gin.Context) {
	id := ctx.Param("id")

	var contact models.Contact
	if err := ctx.ShouldBindJSON(&contact); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var existingContact models.Contact
	result := cc.DB.First(&existingContact, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Contact not found"})
		return
	}

	cc.DB.Model(&existingContact).Updates(contact)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": existingContact})
}

func (cc *ContactController) FindContactById(ctx *gin.Context) {
	id := ctx.Param("id")

	var contact models.Contact
	result := cc.DB.First(&contact, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Contact not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": contact})
}

func (cc *ContactController) FindContacts(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var contacts []models.Contact
	results := cc.DB.Limit(intLimit).Offset(offset).Find(&contacts)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(contacts), "data": contacts})
}

func (cc *ContactController) DeleteContact(ctx *gin.Context) {
	id := ctx.Param("id")

	result := cc.DB.Delete(&models.Contact{}, id)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "Contact not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
