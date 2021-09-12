package database

import (
	"github.com/AnggunPermata/mini-twitter-clone-api/auth"
	"github.com/AnggunPermata/mini-twitter-clone-api/config"
	"github.com/AnggunPermata/mini-twitter-clone-api/models"
)

//CreateUser for saving a new user's data into database
func CreateUser(user models.User) (models.User, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

//UserLogin used to check user data in DB, and if exist it will call CreateToken function to initiate new token.
func UserLogin(email, password string) (models.User, error) {
	var err error
	var user models.User
	if err = config.DB.Where("email=? AND password=?", email, password).First(&user).Error; err != nil {
		return user, err
	}

	user.Token, err = auth.CreateToken(int(user.ID))

	if err != nil {
		return user, err
	}
	return user, nil
}

func GetOneUserById(userId int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", userId).Error; err != nil {
		return user, err
	}
	return user, nil
}
