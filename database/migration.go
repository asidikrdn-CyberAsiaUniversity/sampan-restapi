package database

import (
	"fmt"
	"log"
	"sistem-pengelolaan-bank-sampah/models"
	"sistem-pengelolaan-bank-sampah/pkg/postgres"
)

// migration up
func RunMigration() {
	err := postgres.DB.AutoMigrate(
		&models.Log{},
		&models.MstRole{}, &models.MstUser{},
		&models.MstTrashCategory{}, &models.TrxTrashCustomer{},
	)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Migration failed")
	}

	fmt.Println("Migration up completed successfully")
}

// migration down
func DropMigration() {
	err := postgres.DB.Migrator().DropTable(
		&models.Log{},
		&models.MstRole{}, &models.MstUser{},
		&models.MstTrashCategory{}, &models.TrxTrashCustomer{},
	)

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal("Migration failed")
	}

	fmt.Println("Migration down completed successfully")
}
