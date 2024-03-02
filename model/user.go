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
	Hp       uint   `json:"hp", form:"hp"`
	UserHp   uint   `json:"userhp", form:"userhp"`
	Nama     string `json:"name" form:"name"`
	taskname string `json:"task" form:"task"`
}

type UserModel struct {
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
