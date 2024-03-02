package user

import (
	"mdware/model"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TaskController struct {
	Model model.TaskModel
}
var task []model.Task
func (tc *TaskController) CreateTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		users := c.Get("user").(*jwt.Token)
		claims := users.Claims.(jwt.MapClaims)
		username := claims["username"].(string)

		var taskData model.Task
		if err := c.Bind(&taskData); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err := tc.Model.CreateTask(&taskData).Error();err != nil{
			return err
		}
		return nil
	}
}
