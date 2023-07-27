package handlerTrashCategory

import (
	"sistem-pengelolaan-bank-sampah/dto"
	"sistem-pengelolaan-bank-sampah/models"
	"sistem-pengelolaan-bank-sampah/repositories"
)

type handlerTrashCategory struct {
	TrashCategoryRepository repositories.TrashCategoryRepository
}

func HandlerTrashCategory(trashCategoryRepository repositories.TrashCategoryRepository) *handlerTrashCategory {
	return &handlerTrashCategory{trashCategoryRepository}
}

func convertTrashCategoryResponse(trashCategory *models.MstTrashCategory) *dto.TrashCategoryResponse {
	return &dto.TrashCategoryResponse{
		ID:       trashCategory.ID,
		Category: trashCategory.Category,
		Price:    trashCategory.Price,
	}
}

func convertMultipleTrashCategoryResponse(trashCategory *[]models.MstTrashCategory) *[]dto.TrashCategoryResponse {
	var trashCategoryResponse []dto.TrashCategoryResponse

	for _, tc := range *trashCategory {
		trashCategoryResponse = append(trashCategoryResponse, *convertTrashCategoryResponse(&tc))
	}

	return &trashCategoryResponse
}
