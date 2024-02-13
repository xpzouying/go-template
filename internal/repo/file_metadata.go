package repo

import (
	"context"

	"github.com/xpzouying/go-template/internal/domain"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type fileMetadataModel struct {
	gorm.Model

	FileID   string `gorm:"column:file_id"`
	Filename string `gorm:"column:filename"`
	FileLink string `gorm:"column:file_link"`
}

func (m *fileMetadataModel) TableName() string {
	return "file_metadata"
}

type FileMetadataRepo struct {
	db *gorm.DB
}

func NewFileMetadataRepo(db *gorm.DB) *FileMetadataRepo {
	return &FileMetadataRepo{
		db: db,
	}
}

func (r *FileMetadataRepo) Save(ctx context.Context, metadatas domain.FilesMetadata) error {

	records := make([]fileMetadataModel, 0, len(metadatas))
	for _, metadata := range metadatas {
		records = append(records, fileMetadataModel{
			FileID:   metadata.FileID,
			Filename: metadata.Filename,
			FileLink: metadata.FileLink,
		})
	}

	if err := r.db.WithContext(ctx).Create(records).Error; err != nil {
		return errors.Wrap(err, "failed to save file metadata")
	}

	return nil
}
