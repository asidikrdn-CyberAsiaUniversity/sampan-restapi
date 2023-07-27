package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sistem-pengelolaan-bank-sampah/models"
	"sistem-pengelolaan-bank-sampah/pkg/bcrypt"
	"sistem-pengelolaan-bank-sampah/pkg/postgres"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RunSeeder() {
	// Role
	if postgres.DB.Migrator().HasTable(&models.MstRole{}) {
		newRole := []models.MstRole{}

		newRole = append(newRole, models.MstRole{
			Role: "Superadmin",
		})
		newRole = append(newRole, models.MstRole{
			Role: "Admin",
		})
		newRole = append(newRole, models.MstRole{
			Role: "User",
		})

		for _, role := range newRole {
			errAddRole := postgres.DB.Create(&role).Error
			if errAddRole != nil {
				fmt.Println(errAddRole.Error())
				log.Fatal("Seeding failed")
			}
		}

		fmt.Println("Success seeding master role...")
	}

	// Add Superadmin
	if postgres.DB.Migrator().HasTable(&models.MstUser{}) {
		// check is user table has minimum 1 user
		err := postgres.DB.First(&models.MstUser{}).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// create 1 user
			newUser := models.MstUser{
				ID:              uuid.New(),
				FullName:        "Super Admin",
				Email:           os.Getenv("SUPERADMIN_EMAIL"),
				IsEmailVerified: true,
				LoginAccess:     true,
				RoleID:          1,
			}

			hashPassword, err := bcrypt.HashingPassword(os.Getenv("SUPERADMIN_PASSWORD"))
			if err != nil {
				log.Fatal("Hash password failed")
			}

			newUser.Password = hashPassword

			// insert user to database
			errAddUser := postgres.DB.Create(&newUser).Error
			if errAddUser != nil {
				fmt.Println(errAddUser.Error())
				log.Fatal("Seeding failed")
			}
		}
		fmt.Println("Success seeding super admin...")
	}

	// Trash Category
	if postgres.DB.Migrator().HasTable(&models.MstTrashCategory{}) {
		newTrash := []models.MstTrashCategory{}

		newTrash = append(newTrash, models.MstTrashCategory{
			Category: "KORAN",
			Price:    2000,
		})
		newTrash = append(newTrash, models.MstTrashCategory{
			Category: "BUKU",
			Price:    1000,
		})
		newTrash = append(newTrash, models.MstTrashCategory{
			Category: "DUPLEX",
			Price:    4000,
		})
		newTrash = append(newTrash, models.MstTrashCategory{
			Category: "KARDUS",
			Price:    2000,
		})
		newTrash = append(newTrash, models.MstTrashCategory{
			Category: "BOTOL MINERAL",
			Price:    3000,
		})
		newTrash = append(newTrash, models.MstTrashCategory{
			Category: "GELAS WARNA",
			Price:    4000,
		})
		newTrash = append(newTrash, models.MstTrashCategory{
			Category: "GELAS MINERAL",
			Price:    1500,
		})
		newTrash = append(newTrash, models.MstTrashCategory{
			Category: "BESI",
			Price:    5000,
		})

		for _, trash := range newTrash {
			errAddTrash := postgres.DB.Create(&trash).Error
			if errAddTrash != nil {
				fmt.Println(errAddTrash.Error())
				log.Fatal("Seeding failed")
			}
		}

		fmt.Println("Success seeding master trash...")
	}

	fmt.Println("Seeding completed successfully")
}
