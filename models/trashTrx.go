package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TrxTrashCustomer struct {
	ID              uuid.UUID `gorm:"type:varchar(255);primaryKey"`
	UserID          uuid.UUID `gorm:"type:varchar(255)"`
	TrashID         uint
	TransactionTime time.Time
	Qty             int
	TotalPrice      int
	CreatedBy       uuid.UUID `gorm:"type:varchar(255)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	gorm.DeletedAt
	// parent
	TrashCategory MstTrashCategory `gorm:"foreignKey:TrashID"`
	User          UserResponse     `gorm:"foreignKey:UserID"`
	Creator       CreatorResponse  `gorm:"foreignKey:CreatedBy"`
}
