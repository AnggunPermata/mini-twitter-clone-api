package database

import (
	"github.com/AnggunPermata/mini-twitter-clone-api/config"
	"github.com/AnggunPermata/mini-twitter-clone-api/models"
	"github.com/labstack/echo"
)

func GetTimelineData(userDataTimeline []models.User, c echo.Context) ([]models.Tweet, error) {
	var tweet []models.Tweet
	for i := 0; i < len(userDataTimeline); i++ {
		newId := userDataTimeline[i].ID
		tweetData, err := GetOneTweeetForTimeline(int(newId))
		if err != nil {
			return nil, err
		}
		tweet = append(tweet, tweetData...)
	}
	if err := config.DB.Find(&tweet).Order("created_at DESC").Error; err != nil {
		return tweet, err
	}

	return tweet, nil

}
