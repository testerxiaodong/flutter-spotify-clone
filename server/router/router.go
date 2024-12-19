package router

import (
	"server/controllers"
	"server/middleware"
	"server/models/dao"
	"server/services"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	userService := services.NewUserService(dao.Q)
	userController := controllers.NewUserController(userService)

	// User routes
	authRouter := r.Group("/auth")
	authRouter.POST("/signup", userController.Signup)
	authRouter.POST("/login", userController.Login)
	authRouter.GET("/", middleware.Auth(), userController.UserInfo)
}
