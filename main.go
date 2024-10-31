package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yusufborucu/go-mysql-crud/configs"
	"github.com/yusufborucu/go-mysql-crud/controllers"
	"github.com/yusufborucu/go-mysql-crud/models"
)

func main() {
	r := gin.Default()
	configs.InitDB()
	configs.DB.AutoMigrate(&models.User{})

	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.GetUserById)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run(":8080")
}
