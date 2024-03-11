package user

import (
	"fmt"
	"mdware/helper"
	"mdware/middlewares"
	model "mdware/model/user"

	"log"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model model.UserModel
}

var users []model.User

func (us *UserController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.User
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType, err.Error())
			}
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)
		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest,
					"data yang dikirim kurang sesuai", nil))
		}
		var processInput model.User
		processInput.Hp = input.Hp
		processInput.Nama = input.Nama
		processInput.Email = input.Email
		processInput.Password = input.Password

		err = us.Model.AddUser(processInput)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError,
					"terjadi kesalahan pada sistem", nil))
		}
		return c.JSON(http.StatusCreated,
			helper.ResponseFormat(http.StatusCreated,
				"selamat data sudah terdaftar", nil))
	}
}

func (us *UserController) GetUsersController() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "success get all users",
			"users":    users,
		})
	}
}

func (us *UserController) UpdateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hp = c.Param("hp")
		var input model.User
		err := c.Bind(&input)
		if err != nil {
			log.Println("masalah baca input:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		isFound := us.Model.CekUser(hp)

		if !isFound {
			return c.JSON(http.StatusNotFound,
				helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
		}

		err = us.Model.Update(hp, input)

		if err != nil {
			log.Println("masalah database :", err.Error())
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError, "terjadi kesalahan saat update data", nil))
		}

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "data berhasil di update", nil))
	}
}

func (us *UserController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input model.User
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType, map[string]any{"code": http.StatusUnsupportedMediaType, "message": "format data tidak didukung"})
			}
			return c.JSON(http.StatusBadRequest, map[string]any{"code": http.StatusBadRequest, "message": "data yang dikirimkan tidak sesuai"})
		}
		validate := validator.New(validator.WithRequiredStructEnabled())
		err = validate.Struct(input)

		if err != nil {
			for _, val := range err.(validator.ValidationErrors) {
				fmt.Println(val.Error())
			}
		}
		result, err := us.Model.Login(input.Hp, input.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError,
					"terjadi kesalahan pada sistem", nil))
		}
		token, err := middlewares.GenerateJWT(result.Hp)
		if err != nil {
			return c.JSON(http.StatusInternalServerError,
				helper.ResponseFormat(http.StatusInternalServerError,
					"terjadi kesalahan pada sistem, gagal memproses data", nil))
		}
		var responseData LoginResponse
		responseData.Hp = result.Hp
		responseData.Nama = result.Nama
		responseData.Token = token
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK,
				"selamat anda berhasil login", responseData))
	}
}

func (us *UserController) ListUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		listUser, err := us.Model.GetAllUser()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"code":    http.StatusInternalServerError,
				"message": "terjadi kesalahan pada sistem",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"code":    http.StatusOK,
			"message": "berhasil mendapatkan data",
			"data":    listUser,
		})
	}
}

func (us *UserController) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {
		var hp = c.Param("hp")
		result, err := us.Model.GetProfile(hp)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, map[string]any{
					"code":    http.StatusNotFound,
					"message": "data tidak ditemukan",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"code":    http.StatusInternalServerError,
				"message": "terjadi kesalahan pada sistem",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"code":    http.StatusOK,
			"message": "berhasil mendapatkan data",
			"data":    result,
		})
	}
}
