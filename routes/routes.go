package routes

import (
	"mdware/config"
	task "mdware/controller"
	user "mdware/controller"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ct1 user.UserController, ct2 task.TaskController) {
	jwtConfig := echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}
	c.POST("/users", ct1.Register())
	c.PUT("/users/:hp", ct1.UpdateUserController())
	c.GET("/users", ct1.ListUser(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/users/:hp", ct1.Profile(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/users/:hp", ct1.UpdateUserController(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.POST("/login", ct1.Login())
	taskGroup := c.Group("/tasks", echojwt.WithConfig(jwtConfig))
	taskGroup.POST("", ct2.CreateTask())
	taskGroup.PUT("/:id", ct2.UpdateTask())

}
