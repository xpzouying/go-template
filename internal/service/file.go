package service

import (
	"io"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xpzouying/go-template/api/response"
	"github.com/xpzouying/go-template/internal/domain"
)

func (s *Service) Upload(c *gin.Context) {

	form, err := c.MultipartForm()
	if err != nil {
		logrus.Fatalf("Error when parsing form: %v", err)
		return
	}

	files := form.File["file"]

	fileInfos, err := s.readAllFiles(files)
	if err != nil {
		response.WriteError(c, err.Error())
		return
	}

	ctx := c.Request.Context()
	result, err := s.fileDO.Upload(ctx, fileInfos)
	if err != nil {
		response.WriteError(c, err.Error())
		return
	}

	resp := s.makeUploadResult(result)
	response.WriteSuccess(c, resp)
}

func (s *Service) readAllFiles(files []*multipart.FileHeader) (*domain.FileInfos, error) {
	var fileInfos []*domain.FileInfo

	for _, file := range files {
		f, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer f.Close()

		data, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}

		fileInfos = append(fileInfos, &domain.FileInfo{
			Filename: file.Filename,
			Data:     data,
		})
	}

	return &domain.FileInfos{
		Files: fileInfos,
	}, nil
}

func (s *Service) makeUploadResult(result *domain.FilesUploadResult) []response.UploadResult {
	uploadResults := make([]response.UploadResult, 0, len(result.Results))

	for _, r := range result.Results {
		uploadResults = append(uploadResults, response.UploadResult{
			FileID: r.FileID,
			Link:   r.Link,
		})
	}

	return uploadResults
}
