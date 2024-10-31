package services

import (
	"github.com/yusufborucu/go-mysql-crud/models"
	"github.com/yusufborucu/go-mysql-crud/repositories"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func GetUserById(id uint) (models.User, error) {
	return repositories.GetUserById(id)
}

func CreateUser(user models.User) error {
	return repositories.CreateUser(user)
}

func UpdateUser(user models.User) error {
	return repositories.UpdateUser(user)
}

func DeleteUser(id uint) error {
	return repositories.DeleteUser(id)
}
