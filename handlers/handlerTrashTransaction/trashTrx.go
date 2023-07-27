package handlerTrashTransaction

import (
	"sistem-pengelolaan-bank-sampah/dto"
	"sistem-pengelolaan-bank-sampah/models"
	"sistem-pengelolaan-bank-sampah/repositories"
)

type handlerTrashTransaction struct {
	TrashTransactionRepository repositories.TrashTransactionRepository
}

func HandlerTrashTransaction(trashTransactionRepository repositories.TrashTransactionRepository) *handlerTrashTransaction {
	return &handlerTrashTransaction{trashTransactionRepository}
}

func convertTrashTransactionResponse(trashTransaction *models.TrxTrashCustomer) *dto.TrashTransactionResponse {
	return &dto.TrashTransactionResponse{
		ID:              trashTransaction.ID,
		TransactionTime: trashTransaction.TransactionTime,
		Qty:             trashTransaction.Qty,
		TotalPrice:      uint(trashTransaction.TotalPrice),
		Customer: dto.UserResponse{
			ID:       trashTransaction.User.ID,
			FullName: trashTransaction.User.FullName,
			Email:    trashTransaction.User.Email,
			Phone:    trashTransaction.User.Phone,
			Address:  trashTransaction.User.Address,
			Image:    trashTransaction.User.Image,
			Role: dto.RoleResponse{
				ID:   trashTransaction.User.Role.ID,
				Role: trashTransaction.User.Role.Role,
			},
		},
		Trash: dto.TrashCategoryResponse{
			ID:       trashTransaction.TrashCategory.ID,
			Category: trashTransaction.TrashCategory.Category,
			Price:    trashTransaction.TrashCategory.Price,
		},
		Creator: dto.UserResponse{
			ID:       trashTransaction.Creator.ID,
			FullName: trashTransaction.Creator.FullName,
			Email:    trashTransaction.Creator.Email,
			Phone:    trashTransaction.Creator.Phone,
			Address:  trashTransaction.Creator.Address,
			Image:    trashTransaction.Creator.Image,
			Role: dto.RoleResponse{
				ID:   trashTransaction.Creator.Role.ID,
				Role: trashTransaction.Creator.Role.Role,
			},
		},
	}
}

func convertMultipleTrashTransactionResponse(trashTransaction *[]models.TrxTrashCustomer) *[]dto.TrashTransactionResponse {
	var trxResponse []dto.TrashTransactionResponse

	for _, t := range *trashTransaction {
		trxResponse = append(trxResponse, *convertTrashTransactionResponse(&t))
	}

	return &trxResponse
}
