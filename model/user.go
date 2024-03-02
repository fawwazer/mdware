package model

import (
	"gorm.io/gorm"
)

type User struct {
	Hp       int    `json:"hp" form:"hp"`
	Nama     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Task struct {
	Hp       int    `json:"hp", form:"hp"`
	UserHp   int    `json:"userhp", form:"userhp"`
	Nama     string `json:"name" form:"name"`
	taskname string `json:"task" form:"task"`
}

type UserModel struct {
	Connection *gorm.DB
}

type TaskModel struct {
	Connection *gorm.DB
}

func (um *UserModel) AddUser(newData User) error {
	err := um.Connection.Create(&newData).Error
	if err != nil {
		return err
	}
	return nil
}

func (um *UserModel) CekUser(hp string) bool {
	var data User
	if err := um.Connection.Where("Hp = ?", hp).First(&data).Error; err != nil {
		return false
	}
	return true
}

func (um *UserModel) Update(hp string, data User) error {
	if err := um.Connection.Model(&data).Where("Hp = ?", hp).Update("Name", data.Nama).Update("Password", data.Password).Error; err != nil {
		return err
	}
	return nil
}

func (um *UserModel) GetAllUser() ([]User, error) {
	var result []User

	if err := um.Connection.Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (um *UserModel) GetProfile(hp string) (User, error) {
	var result User
	if err := um.Connection.Where("Hp = ?", hp).First(&result).Error; err != nil {
		return User{}, err
	}
	return result, nil
}

func (um *UserModel) Login(hp int, password string) (User, error) {
	var result User
	if err := um.Connection.Where("Hp = ? AND Password = ?", hp, password).First(&result).Error; err != nil {
		return User{}, err
	}
	return result, nil
}

func (um *UserModel) GetUserIDByUsername(hp int) (User, error) {
	var user User
	if err := um.Connection.Where("hp = ?", hp).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (tm *TaskModel) CreateTask(username string, taskData Task) error {
	var user User
	if err := tm.Connection.Where("username= ?", username).First(&user).Error; err != nil {
		return err
	}
	taskData.UserHp = user.Hp

	if err := tm.Connection.Create(&taskData).Error; err != nil {
		return err
	}
	return nil
}
