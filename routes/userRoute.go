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
	r.GET("/check-auth", middleware.UserAuth(), h.CheckAuth)

	// create new user
	r.POST("/users/staff", middleware.SuperAdminAuth(), middleware.UploadSingleFile(), h.AddStaff)                // add new staff
	r.POST("/users/customer", middleware.AdminAuth(), middleware.UploadSingleFile(), h.AddCustomer)               // add new customer
	r.POST("/users/customer/self-register", middleware.AdminAuth(), middleware.UploadSingleFile(), h.AddCustomer) // self register customer, customer will be have user account to login

	// find/get user
	r.GET("/users/staff", middleware.AdminAuth(), h.FindAllStaff)
	r.GET("/users/customer", middleware.AdminAuth(), h.FindAllCustomer)
	r.GET("/users/:id", middleware.AdminAuth(), h.FindUserByID)
	r.GET("/users/profile", middleware.UserAuth(), h.GetProfile)

	// update user
	r.PATCH("/users/:id", middleware.AdminAuth(), middleware.UploadSingleFile(), h.UpdateUserByID)
	r.PATCH("/users/profile", middleware.UserAuth(), middleware.UploadSingleFile(), h.UpdateProfile)

	// delete user
	r.DELETE("/users/:id", middleware.AdminAuth(), h.DeleteUser)
}
