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

	fmt.Println("Seeding completed successfully")
}
