package router

import (
	"server/controllers"
	"server/models/dao"
	"server/services"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	userService := services.NewUserService(dao.Q)
	userController := controllers.NewUserController(userService)
	v1 := r.Group("/auth")
	v1.POST("/signup", userController.Signup)
}
