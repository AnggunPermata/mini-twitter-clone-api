package route

import (
	"github.com/AnggunPermata/mini-twitter-clone-api/controller"
	"github.com/labstack/echo"
)

func New(e *echo.Echo) {
	// User signup
	e.POST("users/signup", controller.Signup)

	//User Login
	e.POST("users/login", controller.UserLogin)
}
