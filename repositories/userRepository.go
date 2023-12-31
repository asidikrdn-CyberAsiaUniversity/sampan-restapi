package repositories

import (
	"fmt"
	"sistem-pengelolaan-bank-sampah/dto"
	"sistem-pengelolaan-bank-sampah/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.MstUser) (*models.MstUser, error)
	UpdateUser(user *models.MstUser) (*models.MstUser, error)
	DeleteUser(user *models.MstUser) (*models.MstUser, error)
	FindUserByID(id uuid.UUID) (*models.MstUser, error)
	FindAllUsers(limit, offset int, filter dto.UserFilter, searchQuery string) (*[]models.MstUser, int64, error)
	GetUserByEmailOrPhone(searchQuery string) (*models.MstUser, error)
}

func (r *repository) CreateUser(user *models.MstUser) (*models.MstUser, error) {
	err := r.db.Create(user).Error

	return user, err
}

func (r *repository) UpdateUser(user *models.MstUser) (*models.MstUser, error) {
	err := r.db.Model(&user).
		Updates(*user).Error

	return user, err
}

func (r *repository) DeleteUser(user *models.MstUser) (*models.MstUser, error) {
	err := r.db.Delete(user).Error

	return user, err
}

func (r *repository) FindUserByID(id uuid.UUID) (*models.MstUser, error) {
	var users models.MstUser

	err := r.db.Preload("Role").
		Preload("Transaction").
		Preload("TransactionCreated").
		Where("id = ?", id).First(&users).Error
	return &users, err
}

func (r *repository) FindAllUsers(limit, offset int, filter dto.UserFilter, searchQuery string) (*[]models.MstUser, int64, error) {
	var (
		users     []models.MstUser
		totalUser int64
	)

	// create new transaction
	trx := r.db.Session(&gorm.Session{})

	if filter.RoleID != 0 {
		trx = trx.Where("role_id = ?", filter.RoleID)
	} else {
		trx = trx.Where("role_id <> ?", 3) // if filter query not exist, only get superadmin and admin
	}

	// join tables, used for complex searching on relation table
	trx = trx.Joins("JOIN mst_roles ON mst_roles.id = mst_users.role_id")

	if searchQuery != "" {
		trx = trx.Where("full_name LIKE ? OR email LIKE ? OR phone LIKE ? OR address LIKE ? OR mst_roles.role LIKE ?",
			fmt.Sprintf("%%%s%%", searchQuery),
			fmt.Sprintf("%%%s%%", searchQuery),
			fmt.Sprintf("%%%s%%", searchQuery),
			fmt.Sprintf("%%%s%%", searchQuery),
			fmt.Sprintf("%%%s%%", searchQuery))
	}

	// preloading, used for get relation data for results
	trx = trx.Preload("Role").
		Preload("Transaction").
		Preload("TransactionCreated")

	// count transaction result
	trx.Model(&models.MstUser{}).
		Count(&totalUser)

	// set pagination
	err := trx.Limit(limit).
		Offset(offset).
		Find(&users).Error

	return &users, totalUser, err
}

func (r *repository) GetUserByEmailOrPhone(searchQuery string) (*models.MstUser, error) {
	var user models.MstUser

	trx := r.db.Session(&gorm.Session{})

	if searchQuery != "" {
		trx = trx.Where("email = ?", searchQuery).
			Or("phone = ?", searchQuery)
	}

	// preloading, used for get relation data for results
	trx = trx.Preload("Role")

	err := trx.First(&user).Error
	return &user, err
}
