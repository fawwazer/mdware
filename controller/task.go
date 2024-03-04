package user

import (
	"net/http"
	"strconv"

	"mdware/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type TaskController struct {
	Model model.TaskModel
}

func (tc *TaskController) CreateTask() echo.HandlerFunc {
	return func(c echo.Context) error {

		users := c.Get("user").(*jwt.Token)
		claims := users.Claims.(jwt.MapClaims)
		username := claims["username"].(string)

		var taskData model.Task
		if err := c.Bind(&taskData); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		err := tc.Model.CreateTask(username, taskData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusCreated, taskData)
	}
}

func (tc *TaskController) UpdateTask() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskID := c.Param("id")
		taskIDInt, err := strconv.Atoi(taskID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		taskIDUint := uint(taskIDInt)

		var updateTask model.Task
		if err := c.Bind(&updateTask); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		err = tc.Model.UpdateTask(taskIDUint, updateTask)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Task Updated Successfully"})
	}
}
