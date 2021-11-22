package database

import (
	"github.com/jaysonmulwa/go-rest/internal/comment"
	"github.com/jinzhu/gorm"
)

//Migrate DB - migrates the database and create our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
