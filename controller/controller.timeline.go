package controller

import (
	"net/http"
	"strconv"

	"github.com/AnggunPermata/mini-twitter-clone-api/database"
	"github.com/AnggunPermata/mini-twitter-clone-api/models"
	"github.com/labstack/echo"
)

func Timeline(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	// input := models.InputTimeline{}
	// c.Bind(&input)
	followed, err := database.GetFollowedUser(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot get followed user data",
		})
	}
	user, err := database.GetOneUserByIdForTimeline(userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot get followed user data",
		})
	}
	var userDataTimeline []models.User
	userDataTimeline = append(userDataTimeline, user...)

	for i := 0; i < len(followed); i++ {
		newId := followed[i].FollowedUserID
		userData, err := database.GetOneUserByIdForTimeline(int(newId))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "error fetching data (followed user)",
			})
		}
		userDataTimeline = append(userDataTimeline, userData...)
	}

	c.Bind(&userDataTimeline)
	timeline, err := database.GetTimelineData(userDataTimeline, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error setting up timeline",
		})
	}

	var timelineMapping []map[string]interface{}
	for i := 0; i < len(timelineMapping); i++ {
		tweetMaps := map[string]interface{}{
			"Tweet ID": timeline[i].ID,
			"User ID":  timeline[i].UserID,
			"Username": timeline[i].Username,
			"Fullname": timeline[i].FullName,
			"Tweet":    timeline[i].TweetData,
			"Favorite": timeline[i].Favorite,
			"Retweet":  timeline[i].Retweet,
		}
		timelineMapping = append(timelineMapping, tweetMaps)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    timelineMapping,
	})
}
