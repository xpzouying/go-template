package controller

import "go.uber.org/zap"

var (
	logger *zap.SugaredLogger
)

// InitLogger 初始化 logger
func InitLogger(l *zap.SugaredLogger) {
	logger = l
}
