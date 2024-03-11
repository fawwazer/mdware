package main

import (
	"mdware/config"
	"mdware/model/task"
	"mdware/model/user"

	tControll "mdware/controller/task"
	uControll "mdware/controller/user"
	"mdware/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)
	m := user.UserModel{Connection: db}
	c := uControll.UserController{Model: m}
	tm := task.TaskModel{Connection: db}
	tc := tControll.TaskController{Model: tm}
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, c, tc)
	e.Logger.Fatal(e.Start(":8000"))
}
