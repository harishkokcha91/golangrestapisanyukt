package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	BusinessController      controllers.BusinessController
	BusinessRouteController routes.BusinessRouteController
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

	BusinessController = controllers.NewBusinessController(initializers.DB)
	BusinessRouteController = routes.NewBusinessRouteController(BusinessController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	ProfileRouteController.ProfileRoute(router)

	UserRouteController.UserRoute(router)
	PostRouteController.PostRoute(router)
	BusinessRouteController.BusinessRoute(router)

	log.Fatal(server.Run(":" + config.ServerPort))
}
