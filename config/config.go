package config

import (
	"strconv"

	"github.com/AnggunPermata/mini-twitter-clone-api/constant"
	"github.com/AnggunPermata/mini-twitter-clone-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var PORT int

func InitDB() {
	connectionString := constant.Configuration["ConnectionString"]
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitPort() {
	PORT, _ = strconv.Atoi(constant.Configuration["PORT"])
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Tweet{})
	DB.AutoMigrate(&models.Timeline{})
	DB.AutoMigrate(&models.Comment{})
}
