package database

import (
	"github.com/AnggunPermata/mini-twitter-clone-api/config"
	"github.com/AnggunPermata/mini-twitter-clone-api/models"
)

func CreateNewFollower(follow models.Follow) (models.Follow, error) {
	if err := config.DB.Save(&follow).Error; err != nil {
		return follow, err
	}
	return follow, nil
}

func GetFollowedUser(userId int) ([]models.Follow, error) {
	var usersFollowed []models.Follow
	if err := config.DB.Find(&usersFollowed, "id = ? AND follow_status = ?", userId, "followed").Error; err != nil {
		return usersFollowed, err
	}
	return usersFollowed, nil
}
