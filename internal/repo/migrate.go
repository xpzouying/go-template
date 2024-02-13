package repo

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	if err := migrateFileMetadata(db); err != nil {
		return err
	}

	return nil
}

func migrateFileMetadata(db *gorm.DB) error {
	return db.AutoMigrate(&fileMetadataModel{})
}
