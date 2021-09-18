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
