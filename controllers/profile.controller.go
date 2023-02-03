package controllers

import (
	"fmt"
	"net/http"

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

func (ac *ProfileController) GetProfiles(ctx *gin.Context) {
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

func (ac *ProfileController) CreateProfile(ctx *gin.Context) {
	//	{"user_id":"52","profile_for":"child","gender":"male","user_name":"Lakshiv","date_of_birth":"20-12-2000",
	//	"gotra":"kokcha","mother_gotra":"rohilla","mother_tongue":"hindi","native_place":"Narnaul","height":"5.6",
	//	"physique":"Fit","complexion":"fair","weight":"65","blood_group":"A+","hobbies":"Playing cricket",
	//  "family_status":"Middle class","place_of_birth":"Narnaul","time_of_birth":"01:25","have_horoscope":"no",
	//  "manglik":"no","unmarried_brothers":"0","married_brothers":"1","married_sisters":"1","unmarried_sisters":"1",
	//  "marital_status":"single","no_of_children":"0","academic":"Graduate","profession":"Private job","salary":"10lakh",
	//  "company":"XYZ","occupation_address":"Gurgaon","smoke":"no","drink":"no","diet":"veg","other_amusements":"none",
	//  "Address":"Kokcha bhawan moh: peer aga","city":"Narnaul","state":"Haryana","country":"India","pin_code":"123001",
	//  "email_address":"hkokcha@gmail.com","contact_number":"9996898299","father_name":"Pramod","about_self":"Happilly married",
	//  "about_preference":"Graduate required","show_picture_publiclly":"no","created_at":"0001-01-01T00:00:00Z",
	//  "updated_at":"0001-01-01T00:00:00Z"}

	var payload *models.Profile
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, CustomResponse(err.Error(), false, nil))
		return
	}
	result := ac.DB.Create(&payload)

	if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	ctx.JSON(http.StatusConflict, gin.H{"status": "true", "message": "Profile created succeffully"})
}
