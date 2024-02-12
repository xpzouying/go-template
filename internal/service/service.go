package service

import (
	"context"

	"github.com/xpzouying/go-template/api/response"
	"github.com/xpzouying/go-template/internal/domain"
)

// type Service interface {
// 	Status(ctx context.Context) response.StatusResult
// }

type Service struct {
	fileDO domain.FileDO
}

func New(fileDO domain.FileDO) *Service {
	return &Service{
		fileDO: fileDO,
	}
}

func (s *Service) Status(ctx context.Context) response.StatusResult {
	return response.StatusResult{
		Version:   "0.0.1",
		StartTime: "2024-01-01",
	}
}
