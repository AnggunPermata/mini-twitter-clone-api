package config

import (
	"github.com/AnggunPermata/mini-twitter-clone-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var PORT int

func InitDB() {
	connectionString := "root:Teacup21@tcp(localhost:3306)/mini-twitter-api?charset=utf8&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	InitMigrate()
}

func InitPort() {
	PORT = 8080
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Tweet{})
	DB.AutoMigrate(&models.Timeline{})
	DB.AutoMigrate(&models.Comment{})
}
