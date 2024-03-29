package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krishnalagad/gin-gorm-rest/controller"
)

func UserRoute(router *gin.Engine) {
	router.GET("/", controller.UserController)

	userRepo := controller.New()
	router.POST("/api/user/", userRepo.CreateUser)
}
