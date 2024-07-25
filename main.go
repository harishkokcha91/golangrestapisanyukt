package main

import (
	"log"
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wpcodevo/golang-gorm-postgres/docs" // This is required to serve the Swagger docs

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin" // swagger embed files
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"github.com/wpcodevo/golang-gorm-postgres/routes"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	PostController      controllers.PostController
	PostRouteController routes.PostRouteController

	ProfileController      controllers.ProfileController
	ProfileRouteController routes.ProfileRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables ", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	PostController = controllers.NewPostController(initializers.DB)
	PostRouteController = routes.NewRoutePostController(PostController)

	ProfileController = controllers.NewProfileController(initializers.DB)
	ProfileRouteController = routes.NewProfileRouteController(ProfileController)

	server = gin.Default()
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/
func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	// Swagger endpoint
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API group
	router := server.Group("/api")

	// @Summary Health Check
	// @Description Returns a welcome message
	// @Tags health
	// @Accept  json
	// @Produce  json
	// @Success 200 {object} map[string]interface{}
	// @Router /api/healthchecker [get]
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	// Swagger endpoint
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

	AuthRouteController.AuthRoute(router)
	ProfileRouteController.ProfileRoute(router)

	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))

}
