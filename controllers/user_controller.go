package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yusufborucu/go-mysql-crud/models"
	"github.com/yusufborucu/go-mysql-crud/services"
	"github.com/yusufborucu/go-mysql-crud/utils"
)

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to fetch users")
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := services.GetUserById(uint(id))
	if err != nil {
		utils.JSONResponse(c, http.StatusNotFound, "User not found")
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid data")
		return
	}
	if err := services.CreateUser(user); err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to create user")
		return
	}
	utils.JSONResponse(c, http.StatusCreated, "User created successfully")
}

func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "Invalid data")
		return
	}
	user.ID = uint(id)
	if err := services.UpdateUser(user); err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to update user")
		return
	}
	utils.JSONResponse(c, http.StatusOK, "User updated successfully")
}

func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := services.DeleteUser(uint(id)); err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "Failed to delete user")
		return
	}
	utils.JSONResponse(c, http.StatusOK, "User deleted successfully")
}
