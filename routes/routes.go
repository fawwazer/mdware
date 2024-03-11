package routes

import (
	"mdware/config"
	"mdware/controller/task"
	"mdware/controller/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

//func InitRoute(c *echo.Echo, ct1 user.UserController, ct2 task.TaskController) {
//	jwtConfig := echojwt.Config{
//		SigningKey: []byte(config.JWTSECRET),
//	}
//	c.POST("/users", ct1.Register())
//	c.PUT("/users/:hp", ct1.UpdateUserController())
//	c.GET("/users", ct1.ListUser(), echojwt.WithConfig(echojwt.Config{
//		SigningKey: []byte(config.JWTSECRET),
//	}))
//	c.GET("/users/:hp", ct1.Profile(), echojwt.WithConfig(echojwt.Config{
//		SigningKey: []byte(config.JWTSECRET),
//	}))
//	c.PUT("/users/:hp", ct1.UpdateUserController(), echojwt.WithConfig(echojwt.Config{
//		SigningKey: []byte(config.JWTSECRET),
//	}))
//	c.POST("/login", ct1.Login())
//	taskGroup := c.Group("/tasks", echojwt.WithConfig(jwtConfig))
//	taskGroup.POST("", ct2.CreateTask())
//	taskGroup.PUT("/:id", ct2.UpdateTask())

//}

func InitRoute(c *echo.Echo, ct1 user.UserController, tc task.TaskController) {
	userRoute(c, ct1)
	taskRoute(c, tc)
}

func userRoute(c *echo.Echo, ct1 user.UserController) {
	c.POST("/users", ct1.Register())
	c.POST("/login", ct1.Login())
	c.GET("/users", ct1.ListUser(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/profile", ct1.Profile(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	})) // get profile -> butuh penanda khusus
	c.PUT("/users/:hp", ct1.UpdateUserController(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}

func taskRoute(c *echo.Echo, tc task.TaskController) {
	c.POST("/tasks", tc.AddToDo(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))

	c.PUT("/tasks/:taskID", tc.UpdateToDo(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
