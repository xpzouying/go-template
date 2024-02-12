package repo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xpzouying/go-template/internal/domain"
)

func TestSave(t *testing.T) {
	testdir := "/tmp/go-template-test"

	repo := NewLocalFileRepo(WithLocalFilePath(testdir))

	files := []*domain.FileInfo{
		{
			Filename: "test.txt",
			Data:     []byte("test data"),
		},
	}

	result, err := repo.Save(context.Background(), files)
	assert.NoError(t, err)

	assert.Equal(t, len(files), len(result.Results))
	assert.NotEmpty(t, result.Results[0].FileID)
	assert.NotEmpty(t, result.Results[0].Link)
}

func TestSaveOneFile(t *testing.T) {
	repo := NewLocalFileRepo(WithLocalFilePath("/tmp"))

	file := &domain.FileInfo{
		Filename: "test.txt",
		Data:     []byte("test data"),
	}

	fileID, link, err := repo.saveOneFile(file)
	assert.NoError(t, err)

	assert.NotEmpty(t, fileID)
	assert.NotEmpty(t, link)
}
