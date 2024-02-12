package main

import (
	"embed"
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xpzouying/go-template/api"
	"github.com/xpzouying/go-template/internal/domain"
	"github.com/xpzouying/go-template/internal/repo"
	"github.com/xpzouying/go-template/internal/service"
)

//go:embed web/dist/*
var buildFS embed.FS

//go:embed web/dist/index.html
var indexPage []byte

func main() {
	var (
		port       string // http server port
		dev        bool   // 是否为开发模式
		staticPath string // 静态文件路径
	)
	flag.StringVar(&port, "port", ":8080", "port to listen on")
	flag.BoolVar(&dev, "dev", true, "enable dev mode")
	flag.StringVar(&staticPath, "static", "web/dist", "static file path")
	flag.Parse()

	logrus.Infof("start with mode: dev_mode=%v", dev)

	r := gin.New()

	var s *service.Service
	{
		fileDO := newFileDO()
		s = service.New(fileDO)
	}

	staticConfig := api.NewStaticFSConfig(api.WithDevMode(staticPath))

	api.RegisterService(r, s, staticConfig)

	logrus.Infof("listening on port %v", port)
	logrus.Fatal(r.Run(port))
}

func newFileDO() domain.FileDO {
	repo := repo.NewLocalFileRepo(repo.WithLocalFilePath("/tmp"))

	return domain.NewFileDO(repo)
}
