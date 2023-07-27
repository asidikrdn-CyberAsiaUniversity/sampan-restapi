package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MstUser struct {
	ID              uuid.UUID `gorm:"type:varchar(255);primaryKey"`
	FullName        string    `gorm:"type:varchar(255)"`
	Email           string    `gorm:"type:varchar(255);unique"`
	IsEmailVerified bool
	Phone           string `gorm:"type:varchar(255);unique"`
	IsPhoneVerified bool
	Address         string `gorm:"type:text"`
	Password        string `gorm:"type:varchar(255)"`
	LoginAccess     bool
	RoleID          uint
	Image           string `gorm:"type:varchar(255)"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	gorm.DeletedAt
	// parent
	Role MstRole
	// child
	Transaction        []TrxTrashCustomer `gorm:"foreignKey:UserID"`
	TransactionCreated []TrxTrashCustomer `gorm:"foreignKey:CreatedBy"`
}

type UserResponse struct {
	ID       uuid.UUID `gorm:"type:varchar(255);primaryKey"`
	FullName string    `gorm:"type:varchar(255)"`
	Email    string    `gorm:"type:varchar(255)"`
	Phone    string    `gorm:"type:varchar(255)"`
	Address  string    `gorm:"type:text"`
	Password string    `gorm:"type:varchar(255)"`
	Image    string    `gorm:"type:varchar(255)"`
	RoleID   uint
	// parent
	Role MstRole
}

type CreatorResponse struct {
	ID       uuid.UUID `gorm:"type:varchar(255);primaryKey"`
	FullName string    `gorm:"type:varchar(255)"`
	Email    string    `gorm:"type:varchar(255)"`
	Phone    string    `gorm:"type:varchar(255)"`
	Address  string    `gorm:"type:text"`
	Password string    `gorm:"type:varchar(255)"`
	Image    string    `gorm:"type:varchar(255)"`
	RoleID   uint
	// parent
	Role MstRole
}

func (UserResponse) TableName() string {
	return "mst_users"
}

func (CreatorResponse) TableName() string {
	return "mst_users"
}
