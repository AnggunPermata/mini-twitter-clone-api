package controller

import (
	"net/http"
	"strings"

	"github.com/AnggunPermata/mini-twitter-clone-api/database"
	"github.com/AnggunPermata/mini-twitter-clone-api/models"
	"github.com/labstack/echo"
)

func Signup(c echo.Context) error {
	userData := models.User{}
	c.Bind(&userData)
	if len(userData.Username) < 4 || len(userData.Email) < 4 || strings.Contains(userData.Email, ".com") || len(userData.Password) < 6 || len(userData.Gender) > 1 {
		if userData.Gender != "F" && userData.Gender != "M" {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Please follow the rules to Sign up:",
				"rules-1": "Make sure the username has more than 3 characters.",
				"rules-2": "Make sure the email has more than 3 characters, and it is a real email.",
				"rules-3": "Make sure the Password has more than 5 characters.",
				"rules-4": "Gender only have one character, which is F for female or M for male",
			})
		}
	}
	newUser := models.User{}
	newUser.Username = userData.Username
	newUser.Email = userData.Email
	newUser.Password = userData.Password
	newUser.FullName = userData.FullName
	newUser.Gender = userData.Gender
	c.Bind(&newUser)
	user, err := database.CreateUser(newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Cannot Signup",
		})
	}
	showUserData := map[string]interface{}{
		"ID":        user.ID,
		"Name":      user.FullName,
		"Email":     user.Email,
		"Full Name": user.FullName,
		"Gender":    user.Gender,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Succesfully create a new account",
		"data":    showUserData,
	})

}
