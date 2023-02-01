package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"github.com/wpcodevo/golang-gorm-postgres/models"
	"github.com/wpcodevo/golang-gorm-postgres/utils"
	"gorm.io/gorm"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (ac *AuthController) UpdateUser(ctx *gin.Context) {
	var access_token string
	cookie, err := ctx.Cookie("access_token")

	authorizationHeader := ctx.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		access_token = fields[1]
	} else if err == nil {
		access_token = cookie
	}

	if access_token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
		return
	}
	config, _ := initializers.LoadConfig(".")
	sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	fmt.Println(sub)

	var payload *models.SUser

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, CustomResponse(err.Error(), false, nil))
		return
	}
	print(payload)
	id, err := strconv.Atoi(fmt.Sprint(sub))
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusOK, gin.H{"status": "fail", "message": err.Error()})
	} else {
		ac.DB.Model(&models.SUser{}).Where(&models.SUser{ID: id}).Update("name", payload.Name)
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "User Updated succussfully"})
	}
}

func (ac *AuthController) VerifyOtp(ctx *gin.Context) {
	var payload *models.UserSanyukt

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, CustomResponse(err.Error(), false, nil))
		return
	}
	if payload.UserOtp == "" {
		ctx.JSON(http.StatusBadRequest, CustomResponse("Enter OTP", false, nil))
		return
	}

	newUserOtp1 := models.UsersOtp{}
	ac.DB.Where(&models.UsersOtp{UserMobile: payload.UserMobile}).Find(&newUserOtp1)
	fmt.Println(newUserOtp1)
	if newUserOtp1.ID == 0 {
		ctx.JSON(http.StatusBadRequest, CustomResponse("User is not registered", false, nil))
		return
	}
	if newUserOtp1.UserOtp == payload.UserOtp && newUserOtp1.OtpVerified == "0" {
		config, _ := initializers.LoadConfig(".")
		access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, newUserOtp1.UserId, config.AccessTokenPrivateKey)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, newUserOtp1.UserId, config.RefreshTokenPrivateKey)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		print("Update table")
		ac.DB.Model(&models.UsersOtp{}).Where(&models.UsersOtp{UserMobile: payload.UserMobile}).Update("otp_verified", "1")

		ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
		ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
		ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

		ctx.JSON(http.StatusOK, gin.H{"status": "success", "refresh_token": refresh_token, "access_token": access_token})
		return
	} else {
		ctx.JSON(http.StatusBadRequest, CustomResponse("OTP Not Verified", false, nil))
		return
	}

}

func (ac *AuthController) GenerateOtp(ctx *gin.Context) {
	var payload *models.UserSanyukt

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, CustomResponse(err.Error(), false, nil))
		return
	}
	if payload.UserMobile == "" {
		ctx.JSON(http.StatusBadRequest, CustomResponse("user mobile cannot be empty", false, nil))
		return
	}

	now := time.Now()
	newUser := models.SUser{
		Mobile:    payload.UserMobile,
		Role:      "user",
		Verified:  true,
		Photo:     "",
		Provider:  "local",
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ac.DB.Create(&newUser)
	payload.Id = strconv.Itoa(newUser.ID)

	if result.Error != nil && strings.Contains(result.Error.Error(), "Duplicate") {
		ac.createOtpAndUpdateTable(payload, ctx)
		//ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}
	ac.createOtpAndUpdateTable(payload, ctx)
	// userResponse := &models.SUserResponse{
	// 	ID:        newUser.ID,
	// 	Name:      newUser.Name,
	// 	Mobile:    newUser.Mobile,
	// 	Photo:     newUser.Photo,
	// 	Role:      newUser.Role,
	// 	Provider:  newUser.Provider,
	// 	CreatedAt: newUser.CreatedAt,
	// 	UpdatedAt: newUser.UpdatedAt,
	// }
	// ctx.JSON(http.StatusCreated, CustomResponse("User created succesfully", true, userResponse))

}

func (ac *AuthController) createOtpAndUpdateTable(payload *models.UserSanyukt, ctx *gin.Context) {
	print("genrate otp for mobile " + payload.UserMobile)
	fmt.Println(payload)
	now := time.Now()
	newUserOtp := models.UsersOtp{
		UserMobile:  payload.UserMobile,
		UserOtp:     "877938",
		OtpVerified: "0",
		UserId:      payload.Id,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	newUserOtp1 := models.UsersOtp{}

	if ac.DB.Model(&newUserOtp1).Where("user_mobile=?", newUserOtp.UserMobile).Updates(&models.UsersOtp{UserOtp: "123456", OtpVerified: "0"}).RowsAffected == 0 {
		fmt.Println("User already available")
		fmt.Println(newUserOtp)
		result := ac.DB.Create(&newUserOtp)

		if result.Error != nil {
			ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
			return
		}
		userResponse := &models.UsersOtpResponse{
			UserOtp: newUserOtp.UserOtp,
		}
		ctx.JSON(http.StatusCreated, CustomResponse("Otp has been sent to given mobile number", true, userResponse))
	} else {
		userResponse := &models.UsersOtpResponse{
			UserOtp: "123456",
		}
		ctx.JSON(http.StatusCreated, CustomResponse("Otp has been sent to given mobile number", true, userResponse))

	}

}

// SignUp User
func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *models.SignUpInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Passwords do not match"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	now := time.Now()
	newUser := models.User{
		Name:      payload.Name,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		Role:      "user",
		Verified:  true,
		Photo:     payload.Photo,
		Provider:  "local",
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "User with that email already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	userResponse := &models.UserResponse{
		ID:        newUser.ID,
		Name:      newUser.Name,
		Email:     newUser.Email,
		Photo:     newUser.Photo,
		Role:      newUser.Role,
		Provider:  newUser.Provider,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

func (ac *AuthController) SignInUser(ctx *gin.Context) {
	var payload *models.SignInInput

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or Password"})
		return
	}

	config, _ := initializers.LoadConfig(".")

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refresh_token, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

// Refresh Access Token
func (ac *AuthController) RefreshAccessToken(ctx *gin.Context) {
	message := "could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	config, _ := initializers.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := ac.DB.First(&user, "id = ?", fmt.Sprint(sub))
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", access_token, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": access_token})
}

func (ac *AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
