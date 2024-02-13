package repo

import (
	"context"
	"os"
	"path/filepath"

	"github.com/xpzouying/go-template/internal/domain"
	"github.com/xpzouying/go-template/internal/pkg/myuuid"
)

type (
	LocalFileRepo struct {
		Path string
	}

	LocalFileConfig struct {
		Path string
	}

	LocalFileConfigOption func(*LocalFileConfig)
)

func newDefLocalFileConfig() *LocalFileConfig {
	return &LocalFileConfig{
		Path: "/tmp",
	}
}

func WithLocalFilePath(path string) LocalFileConfigOption {
	return func(cfg *LocalFileConfig) {
		cfg.Path = path
	}
}

func NewLocalFileRepo(options ...LocalFileConfigOption) *LocalFileRepo {
	cfg := newDefLocalFileConfig()

	for _, option := range options {
		option(cfg)
	}

	return &LocalFileRepo{
		Path: cfg.Path,
	}
}

func (r *LocalFileRepo) Save(_ context.Context, files domain.FileInfos) (domain.FilesUploadResult, error) {

	results := make(domain.FilesUploadResult, 0, len(files))

	for _, file := range files {

		fileID, link, err := r.saveOneFile(file)
		if err != nil {
			return nil, err
		}

		results = append(results, domain.FileUploadResult{
			FileID: fileID,
			Link:   link,
		})
	}

	return results, nil
}

func (r *LocalFileRepo) saveOneFile(file *domain.FileInfo) (fileID string, link string, err error) {

	basename := filepath.Base(file.Filename)
	ext := filepath.Ext(basename)
	fileID = myuuid.New()
	link = fileID + ext
	savePath := filepath.Join(r.Path, link)

	if err := r.writeFile(savePath, file.Data); err != nil {
		return "", "", err
	}

	return fileID, link, nil
}

func (r *LocalFileRepo) writeFile(savePath string, data []byte) error {

	dir := filepath.Dir(savePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(savePath, data, 0644); err != nil {
		return err
	}

	return nil
}
