package routes

import (
	"sistem-pengelolaan-bank-sampah/handlers/handlerUser"
	"sistem-pengelolaan-bank-sampah/pkg/middleware"
	"sistem-pengelolaan-bank-sampah/pkg/postgres"
	"sistem-pengelolaan-bank-sampah/repositories"

	"github.com/gin-gonic/gin"
)

func User(r *gin.RouterGroup) {
	userRepository := repositories.MakeRepository(postgres.DB)
	h := handlerUser.HandlerUser(userRepository)

	// login
	r.POST("/login", h.Login)

	// create new user
	r.POST("/users", middleware.AdminAuth(), middleware.UploadSingleFile(), h.CreateUser)

	// find/get user
	r.GET("/users", middleware.AdminAuth(), h.FindAllUsers)
	r.GET("/users/:id", middleware.AdminAuth(), h.FindUserByID)
	r.GET("/users/profile", middleware.UserAuth(), h.GetProfile)

	// update user
	r.POST("/users/:id", middleware.AdminAuth(), middleware.UploadSingleFile(), h.UpdateUserByID)
	r.POST("/users/profile", middleware.UserAuth(), middleware.UploadSingleFile(), h.UpdateProfile)

	// delete user
	r.DELETE("/users/:id", middleware.SuperAdminAuth(), h.DeleteUser)
}
