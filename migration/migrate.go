package migration

import (
	"mdware/model"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// Check if the 'hp' column already exists in the 'users' table
	hasHPColumn := db.Migrator().HasColumn(&model.User{}, "hp")
	if !hasHPColumn {
		// If the 'hp' column doesn't exist, add it
		if err := db.Migrator().AddColumn(&model.User{}, "hp"); err != nil {
			return err
		}
	}

	// Your other migration steps go here

	return nil
}
