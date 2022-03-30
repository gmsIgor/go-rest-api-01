package database

import (
	"github.com/lil-newty/go-rest-api-01/internal/comment"
	"gorm.io/gorm"
)

// MigrateDB - migrates our database and createes our comment table
func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&comment.Comment{}); err != nil {
		return err
	}

	return nil
}
