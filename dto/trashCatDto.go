package dto

type CreateTrashCategoryRequest struct {
	Category string `json:"category" binding:"required"`
	Price    int    `json:"price" binding:"required"`
}

type UpdateTrashCategoryRequest struct {
	Category string `json:"category"`
	Price    int    `json:"price"`
}

type TrashCategoryResponse struct {
	ID       uint   `json:"id"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

type TrashCategoryFilter struct {
	Category string
	Price    uint
}
