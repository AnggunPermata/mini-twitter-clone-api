package controller

import (
	"net/http"
	"strconv"

	"github.com/AnggunPermata/mini-twitter-clone-api/database"
	"github.com/AnggunPermata/mini-twitter-clone-api/models"
	"github.com/labstack/echo"
)

func NewTweet(c echo.Context) error {
	//to get user id using param
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	if AuthorizedUser(c) == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized user")
	}

	var tweetData string
	data := models.Tweet{}
	c.Bind(&data)

	//tweetData has maximal length of characters, which is 300.
	if len(data.TweetData) > 300 {
		tweetData = data.TweetData[:301]
	} else {
		tweetData = data.TweetData
	}

	newTweet := models.Tweet{}
	userData, err := database.GetOneUserById(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "user is not available.",
		})
	}

	newTweet.UserID = userData.ID
	newTweet.Username = userData.Username
	newTweet.FullName = userData.FullName
	newTweet.TweetData = tweetData
	newTweet.Favorite = 0
	newTweet.Retweet = 0
	c.Bind(&newTweet)

	//save the new tweet data into database
	tweet, err := database.CreateNewTweet(newTweet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Cannot post a new tweet.",
		})
	}

	showTweetData := map[string]interface{}{
		"TweetID":   tweet.ID,
		"UserID":    tweet.UserID,
		"Username":  tweet.Username,
		"FullName":  tweet.FullName,
		"TweetData": tweet.TweetData,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success creating a new tweet.",
		"data":    showTweetData,
	})
}
