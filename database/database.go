package database

import (
	"fmt"
	"strconv"

	"github.com/chiragthapa777/wishlist-api/config"
	"github.com/chiragthapa777/wishlist-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// ConnectDB connect to db
func ConnectDB() {
	var config = config.GetConfig()
	p := config.DatabasePort
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic("failed to parse database port")
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.DatabaseHost, port, config.DatabaseUser, config.DatabasePassword, config.DatabaseName)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db = database

	fmt.Println("Connection Opened to Database")
	db.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
}

func GetDatabase() *gorm.DB {
	return db
}
