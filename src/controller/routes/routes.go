package routes

import (
	"github.com/andregtalima/CRUD-Golang/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createtUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpadteUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}