package repositories

import (
	"fmt"
	"sistem-pengelolaan-bank-sampah/dto"
	"sistem-pengelolaan-bank-sampah/models"

	"gorm.io/gorm"
)

type TrashCategoryRepository interface {
	FindAllTrashCategory(limit, offset int, filterQuery dto.TrashCategoryFilter, searchQuery string) (*[]models.MstTrashCategory, int64, error)
	FindTrashCategoryByID(trashID uint) (*models.MstTrashCategory, error)
	CreateTrashCategory(trash *models.MstTrashCategory) (*models.MstTrashCategory, error)
	UpdateTrashCategory(trash *models.MstTrashCategory) (*models.MstTrashCategory, error)
	DeleteTrashCategory(trash *models.MstTrashCategory) (*models.MstTrashCategory, error)
}

func (r *repository) FindAllTrashCategory(limit, offset int, filter dto.TrashCategoryFilter, searchQuery string) (*[]models.MstTrashCategory, int64, error) {
	var (
		trash      []models.MstTrashCategory
		totalTrash int64
	)

	trx := r.db.Session(&gorm.Session{})

	if filter.Category != "" {
		trx = trx.Where("category = ?", filter.Category)
	}

	if filter.Price != 0 {
		trx = trx.Where("price = ?", filter.Price)
	}

	if searchQuery != "" {
		trx = trx.Where("category LIKE ? OR price LIKE ?",
			fmt.Sprintf("%%%s%%", searchQuery), // category
			fmt.Sprintf("%%%s%%", searchQuery)) // price
	}

	trx.Model(&models.MstTrashCategory{}).
		Count(&totalTrash)

	err := trx.Limit(limit).
		Offset(offset).
		Find(&trash).Error

	return &trash, totalTrash, err
}

func (r *repository) FindTrashCategoryByID(trashCategoryID uint) (*models.MstTrashCategory, error) {
	var trash models.MstTrashCategory
	err := r.db.Where("id = ?", trashCategoryID).
		First(&trash).Error

	return &trash, err
}

func (r *repository) CreateTrashCategory(trashCategory *models.MstTrashCategory) (*models.MstTrashCategory, error) {
	err := r.db.Create(trashCategory).Error

	return trashCategory, err
}

func (r *repository) UpdateTrashCategory(trashCategory *models.MstTrashCategory) (*models.MstTrashCategory, error) {
	err := r.db.Model(&trashCategory).
		Updates(*trashCategory).Error

	return trashCategory, err
}

func (r *repository) DeleteTrashCategory(trashCategory *models.MstTrashCategory) (*models.MstTrashCategory, error) {
	err := r.db.Delete(trashCategory).Error

	return trashCategory, err
}
