package database

import (
	"fmt"

	"github.com/koficypher/simple_server/models"
	"gorm.io/gorm"
)

// MigrateModels - automigrate models
func MigrateModels(db *gorm.DB) error {
	fmt.Println("Migrating models ....")
	var model models.Article

	if err := db.AutoMigrate(&model); err != nil {
		return err
	}

	fmt.Println("Done Migrating models ...")

	return nil
}
