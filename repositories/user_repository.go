package repositories

import (
	"github.com/yusufborucu/go-mysql-crud/configs"
	"github.com/yusufborucu/go-mysql-crud/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := configs.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id uint) (models.User, error) {
	var user models.User
	if err := configs.DB.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user models.User) error {
	return configs.DB.Create(&user).Error
}

func UpdateUser(user models.User) error {
	return configs.DB.Model(&user).Omit("created_at").Updates(user).Error
}

func DeleteUser(id uint) error {
	return configs.DB.Delete(&models.User{}, id).Error
}
