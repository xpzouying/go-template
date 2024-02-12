package domain

import (
	"context"
)

type (
	FileInfo struct {
		Filename string
		Data     []byte
	}

	FileInfos struct {
		Files []*FileInfo
	}

	FileUploadResult struct {
		// FileID 为每一个上传的文件生成的唯一标识
		FileID string

		// Link 每一个上传的文件的下载链接
		Link string
	}

	FilesUploadResult struct {
		Results []FileUploadResult
	}
)

type Uploader interface {
	Upload(ctx context.Context, fileInfos *FileInfos) (*FilesUploadResult, error)
}

type FileDO interface {
	Uploader
}

type FileRepo interface {
	Save(ctx context.Context, files []*FileInfo) (*FilesUploadResult, error)
}

type fileDO struct {
	fileRepo FileRepo
}

func NewFileDO(fileRepo FileRepo) FileDO {
	return &fileDO{
		fileRepo: fileRepo,
	}
}

func (l *fileDO) Upload(ctx context.Context, fileInfos *FileInfos) (*FilesUploadResult, error) {
	files := fileInfos.Files

	if len(files) == 0 {
		return &FilesUploadResult{[]FileUploadResult{}}, nil
	}

	return l.fileRepo.Save(ctx, files)
}
