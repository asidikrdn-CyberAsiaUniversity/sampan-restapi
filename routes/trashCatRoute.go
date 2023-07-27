package routes

import (
	"sistem-pengelolaan-bank-sampah/handlers/handlerTrashCategory"
	"sistem-pengelolaan-bank-sampah/pkg/middleware"
	"sistem-pengelolaan-bank-sampah/pkg/postgres"
	"sistem-pengelolaan-bank-sampah/repositories"

	"github.com/gin-gonic/gin"
)

func TrashCategory(r *gin.RouterGroup) {
	trashRepository := repositories.MakeRepository(postgres.DB)

	h := handlerTrashCategory.HandlerTrashCategory(trashRepository)

	// find/get
	r.GET("/trash/categories", middleware.UserAuth(), h.FindAllTrashCategory)
	r.GET("/trash/categories/:id", middleware.UserAuth(), h.FindTrashCategoryByID)

	// create
	r.POST("/trash/categories", middleware.AdminAuth(), h.CreateTrashCategory)

	// update
	r.PATCH("/trash/categories/:id", middleware.AdminAuth(), h.UpdateTrashCategory)

	// delete
	r.DELETE("/trash/categories/:id", middleware.AdminAuth(), h.DeleteTrashCategory)
}
