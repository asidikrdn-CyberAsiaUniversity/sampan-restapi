package routes

import (
	"sistem-pengelolaan-bank-sampah/handlers/handlerTrashTransaction"
	"sistem-pengelolaan-bank-sampah/pkg/middleware"
	"sistem-pengelolaan-bank-sampah/pkg/postgres"
	"sistem-pengelolaan-bank-sampah/repositories"

	"github.com/gin-gonic/gin"
)

func TrashTransaction(r *gin.RouterGroup) {
	trashTransactionRepository := repositories.MakeRepository(postgres.DB)
	h := handlerTrashTransaction.HandlerTrashTransaction(trashTransactionRepository)

	// Get Summary Transaction
	r.GET("/dashboard", middleware.AdminAuth(), h.GetTrxSummary)

	// find All Transaction
	r.GET("/trash/transactions", middleware.AdminAuth(), h.FindAllTrashTransaction)

	// findOne
	r.GET("/trash/transactions/:id", middleware.UserAuth(), h.FindTrashTransactionByID)

	// create
	r.POST("/trash/transactions", middleware.AdminAuth(), h.CreateTrashTransaction)

	// update
	r.PATCH("/trash/transactions/:id", middleware.AdminAuth(), h.UpdateTrashTransaction)

	// delete
	r.DELETE("/trash/transactions/:id", middleware.AdminAuth(), h.DeleteTrashTransaction)
}
