package main

import (
	"mdware/config"
	task "mdware/controller"
	user "mdware/controller"
	"mdware/routes"

	"mdware/model"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)
	m := model.UserModel{Connection: db}
	c := user.UserController{Model: m}
	t := task.TaskController{}
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, c, t)
	e.Logger.Fatal(e.Start(":8000"))
}
