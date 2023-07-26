package routes

import (
	"sistem-pengelolaan-bank-sampah/handlers/handlerRole"
	"sistem-pengelolaan-bank-sampah/pkg/middleware"
	"sistem-pengelolaan-bank-sampah/pkg/postgres"
	"sistem-pengelolaan-bank-sampah/repositories"

	"github.com/gin-gonic/gin"
)

func Role(r *gin.RouterGroup) {
	roleRepository := repositories.MakeRepository(postgres.DB)
	h := handlerRole.HandlerRole(roleRepository)

	// find all role
	r.GET("/roles", middleware.UserAuth(), h.FindAllRole)

	// find role by id
	r.GET("/roles/:id", middleware.UserAuth(), h.FindRoleByID)
}
