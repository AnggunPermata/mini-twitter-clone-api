package route

import (
	"github.com/AnggunPermata/mini-twitter-clone-api/constant"
	"github.com/AnggunPermata/mini-twitter-clone-api/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {
	// User signup
	e.POST("users/signup", controller.Signup)

	//User Login
	e.POST("users/login", controller.UserLogin)

	//JWT Authorization
	eJwt := e.Group("")
	eJwt.Use(middleware.JWT([]byte(constant.SECRET_JWT)))
	//User posting new tweet
	eJwt.POST("users/:user_id/new_tweet", controller.NewTweet)
}
