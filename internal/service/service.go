package service

import (
	"context"

	"github.com/xpzouying/go-template/api/response"
)

type Service interface {
	Status(ctx context.Context) response.StatusResult
}

type service struct{}

func New() Service {
	return &service{}
}

func (s *service) Status(ctx context.Context) response.StatusResult {
	return response.StatusResult{
		Version:   "0.0.1",
		StartTime: "2024-01-01",
	}
}
