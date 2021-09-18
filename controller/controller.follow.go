package controller

import (
	"net/http"
	"strconv"

	"github.com/AnggunPermata/mini-twitter-clone-api/database"
	"github.com/AnggunPermata/mini-twitter-clone-api/models"
	"github.com/labstack/echo"
)

func FollowAUser(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	//check Authorization
	if AuthorizedUser(c) == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized user")
	}

	userId2, err2 := strconv.Atoi(c.Param("another_user_id"))
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	follower, err3 := database.GetOneUserById(userId)
	if err3 != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error getting user data",
		})
	}

	input := models.Follow{}
	c.Bind(&input)
	followData := models.Follow{}
	followData.FollowedUserID = uint(userId2)

	followData.UserID = follower.ID
	followData.FullName = follower.FullName
	followData.Username = follower.Username
	followData.FollowStatus = "follow"

	c.Bind(&followData)

	follow, err4 := database.CreateNewFollower(followData)
	if err4 != nil {
		if err3 != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Cannot follow user",
			})
		}
	}
	showFollowingStatus := map[string]interface{}{
		"UserID":          follow.UserID,
		"Followed UserID": follow.FollowedUserID,
		"Follow Status":   follow.FollowStatus,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "following succes",
		"data":    showFollowingStatus,
	})

}
