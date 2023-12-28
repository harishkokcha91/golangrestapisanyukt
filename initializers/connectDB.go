package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error
	// dsn := fmt.Sprintf("host=%s user=%s password=%s
	// dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
	// config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	// 	POSTGRES_HOST=localhost
	// POSTGRES_USER=postgres
	// POSTGRES_PASSWORD=admin
	// POSTGRES_DB=golang-gorm
	// POSTGRES_PORT=5432
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "admin", "golang-gorm")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DBUserName, config.DBUserPassword, config.DBHost, config.DBPort, config.DBName)
	fmt.Println(dsn)
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}
