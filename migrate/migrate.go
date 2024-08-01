package main

import (
	"fmt"
	"log"

	"github.com/wpcodevo/golang-gorm-postgres/initializers"
	"github.com/wpcodevo/golang-gorm-postgres/models"
)

var AllModels = []interface{}{
	&models.User{},
	&models.SUser{},
	&models.UsersOtp{},
	&models.Profile{},
	&models.FavProfile{},
	&models.Business{},
	&models.BusinessHour{},
	&models.Category{},
	&models.BusinessReview{},
	&models.Post{},
	&models.Event{},
}

func init() {
	config, err := initializers.LoadConfig("../.")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables ", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(AllModels...)
	fmt.Println("👍 Migration complete")
}
