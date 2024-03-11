package user

import (
	"mdware/model/task"

	"gorm.io/gorm"
)

type User struct {
	Hp       string `gorm:"type:varchar(13);primaryKey"`
	Nama     string
	Email    string
	Password string
	tasks    []task.Task `gorm:"foreignKey:Pemilik;references:Hp"`
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

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		Connection: db,
	}
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

func (um *UserModel) Login(hp string, password string) (User, error) {
	var result User
	if err := um.Connection.Where("Hp = ? AND Password = ?", hp, password).First(&result).Error; err != nil {
		return User{}, err
	}
	return result, nil
}
