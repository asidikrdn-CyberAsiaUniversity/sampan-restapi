package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateTrashTransactionRequest struct {
	UserID  string `json:"userId" binding:"required"`
	TrashID uint   `json:"trashId" binding:"required"`
	Qty     int    `json:"qty" binding:"required"`
}

type UpdateTrashTransactionRequest struct {
	UserID  string `json:"userId"`
	TrashID uint   `json:"trashId"`
	Qty     int    `json:"qty"`
}

type TrashTransactionResponse struct {
	ID              uuid.UUID             `json:"id"`
	TransactionTime time.Time             `json:"transactionTime"`
	Qty             int                   `json:"qty"`
	TotalPrice      uint                  `json:"totalPrice"`
	Customer        UserResponse          `json:"customer"`
	Trash           TrashCategoryResponse `json:"trash"`
	Creator         UserResponse          `json:"createdBy"`
}

type TrashTransactionFilter struct {
	UserID  string
	TrashID uint
}
