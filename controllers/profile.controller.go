package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"gorm.io/gorm"
)

type ProfileController struct {
	DB *gorm.DB
}

func NewProfileController(DB *gorm.DB) ProfileController {
	return ProfileController{DB}
}

func (ac *ProfileController) GetProfile(ctx *gin.Context) {
	var payload *models.FavProfile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, CustomResponse(err.Error(), false, nil))
		return
	}
	var profiles models.Profile
	ac.DB.Model(&models.Profile{}).Where("id=?", payload.ProfileId).Find(&profiles)

	fmt.Println(profiles)

	ctx.JSON(http.StatusBadRequest, CustomResponse("found", true, profiles))
}

func (ac *ProfileController) GetMyProfiles(ctx *gin.Context) {
	var payload *models.Profile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, CustomResponse(err.Error(), false, nil))
		return
	}
	var profiles []models.Profile
	ac.DB.Model(&models.Profile{}).Where("user_id=?", payload.UserId).Find(&profiles)

	fmt.Println(profiles)

	ctx.JSON(http.StatusBadRequest, CustomResponse("found", true, profiles))
}

func (ac *ProfileController) GetProfilesPost(ctx *gin.Context) {
	var payload *models.Profile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, CustomResponse(err.Error(), false, nil))
		return
	}
	var profiles []models.Profile
	ac.DB.Model(&models.Profile{}).Where("gender=?", payload.Gender).Find(&profiles)

	fmt.Println(profiles)

	ctx.JSON(http.StatusBadRequest, CustomResponse("found", true, profiles))
}

func (ac *ProfileController) CreateFavProfiles(ctx *gin.Context) {
	var payload *models.FavProfile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, CustomResponse(err.Error(), false, nil))
		return
	}
	result := ac.DB.Create(&payload)

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	ctx.JSON(http.StatusConflict, gin.H{"status": "true", "message": "Profile created Fav"})
}

func (ac *ProfileController) GetFavProfiles(ctx *gin.Context) {
	var payload *models.FavProfile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, CustomResponse(err.Error(), false, nil))
		return
	}
	var profiles []models.Profile
	ac.DB.Model(&models.Profile{}).Joins("left join fav_profiles on profiles.id =fav_profiles.profile_id").Where("fav_profiles.user_id=?", payload.UserId).Scan(&profiles)

	ctx.JSON(http.StatusBadRequest, CustomResponse("found", true, profiles))
}

func (pc *ProfileController) CreateProfile(ctx *gin.Context) {
	var payload models.Profile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	payload.CreatedAt = time.Now()
	payload.UpdatedAt = time.Now()

	if result := pc.DB.Create(&payload); result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": payload})
}

func (pc *ProfileController) GetProfiles(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	limit := ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var profiles []models.Profile
	result := pc.DB.Limit(intLimit).Offset(offset).Find(&profiles)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(profiles), "data": profiles})
}

func (pc *ProfileController) GetProfileByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var profile models.Profile
	result := pc.DB.First(&profile, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Profile not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": profile})
}

func (pc *ProfileController) UpdateProfile(ctx *gin.Context) {
	id := ctx.Param("id")

	var payload models.Profile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	var profile models.Profile
	result := pc.DB.First(&profile, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Profile not found"})
		return
	}

	payload.UpdatedAt = time.Now()
	pc.DB.Model(&profile).Updates(payload)
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": profile})
}

func (pc *ProfileController) DeleteProfile(ctx *gin.Context) {
	id := ctx.Param("id")

	result := pc.DB.Delete(&models.Profile{}, id)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "Profile not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
