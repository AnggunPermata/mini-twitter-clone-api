package main

import (
	"fmt"

	"github.com/AnggunPermata/mini-twitter-clone-api/auth"
	"github.com/AnggunPermata/mini-twitter-clone-api/config"
	"github.com/AnggunPermata/mini-twitter-clone-api/route"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	config.InitDB()
	config.InitPort()
	auth.LogMiddlewares((e))
	route.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.PORT)))
}
