package domain

import "context"

type FileRepo interface {
	Save(ctx context.Context, files FileInfos) (FilesUploadResult, error)
}

type (
	FileMetadata struct {
		FileID   string
		Filename string
		FileLink string
	}

	FilesMetadata []FileMetadata
)

type FileMetadataRepo interface {
	Save(ctx context.Context, metadatas FilesMetadata) error
}
