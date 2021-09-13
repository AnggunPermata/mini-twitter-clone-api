package database

import (
	"github.com/AnggunPermata/mini-twitter-clone-api/config"
	"github.com/AnggunPermata/mini-twitter-clone-api/models"
)

func CreateNewTweet(tweet models.Tweet) (models.Tweet, error) {
	if err := config.DB.Save(&tweet).Error; err != nil {
		return tweet, err
	}
	return tweet, nil
}
