package domain

import (
	"context"
)

type (
	FileInfo struct {
		Filename string
		Data     []byte
	}

	FileInfos []*FileInfo

	FileUploadResult struct {
		// FileID 为每一个上传的文件生成的唯一标识
		FileID string

		// Link 每一个上传的文件的下载链接
		Link string
	}

	FilesUploadResult []FileUploadResult
)

type FileDO interface {
	Upload(ctx context.Context, fileInfos FileInfos) (FilesUploadResult, error)
}

type fileDO struct {
	fileRepo FileRepo

	fileMetadataRepo FileMetadataRepo
}

func NewFileDO(
	fileRepo FileRepo,
	fileMetadataRepo FileMetadataRepo,
) FileDO {

	return &fileDO{
		fileRepo:         fileRepo,
		fileMetadataRepo: fileMetadataRepo,
	}
}

func (l *fileDO) Upload(ctx context.Context, fileInfos FileInfos) (FilesUploadResult, error) {

	if len(fileInfos) == 0 {
		return FilesUploadResult{}, nil
	}

	filesUploadResult, err := l.fileRepo.Save(ctx, fileInfos)
	if err != nil {
		return nil, err
	}

	metadatas := make([]FileMetadata, 0, len(fileInfos))
	for i := 0; i < len(fileInfos); i++ {

		metadatas = append(metadatas, FileMetadata{
			Filename: fileInfos[i].Filename,
			FileID:   filesUploadResult[i].FileID,
			FileLink: filesUploadResult[i].Link,
		})
	}

	if err := l.fileMetadataRepo.Save(ctx, metadatas); err != nil {
		return nil, err
	}

	return filesUploadResult, nil
}
