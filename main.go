package main

import (
	"embed"
	"flag"

	"github.com/xpzouying/go-template/api"
	"github.com/xpzouying/go-template/internal/domain"
	"github.com/xpzouying/go-template/internal/repo"
	"github.com/xpzouying/go-template/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

//go:embed web/dist/*
var buildFS embed.FS

//go:embed web/dist/index.html
var indexPage []byte

var (
	port       string // http server port
	dev        bool   // 是否为开发模式
	staticPath string // 静态文件路径
	dbPath     string // 数据库路径
)

func main() {
	flag.StringVar(&port, "port", ":8080", "port to listen on")
	flag.BoolVar(&dev, "dev", true, "enable dev mode")
	flag.StringVar(&staticPath, "static", "web/dist", "static file path")
	flag.StringVar(&dbPath, "db", "/tmp/gorm.db", "database file path")
	flag.Parse()

	logrus.Infof("start with mode: dev_mode=%v", dev)

	db := newDB()

	var s *service.Service
	{
		fileDO := newFileDO(db)
		s = service.New(fileDO)
	}

	staticConfig := api.NewStaticFSConfig(api.WithDevMode(staticPath))

	r := gin.New()
	api.RegisterService(r, s, staticConfig)

	logrus.Infof("listening on port %v", port)
	logrus.Fatal(r.Run(port))
}

func newDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logrus.Fatalf("failed to connect database: %v", err)
	}

	if err = repo.Migrate(db); err != nil {
		logrus.Fatalf("failed to migrate database: %v", err)
	}

	return db
}

func newFileDO(db *gorm.DB) domain.FileDO {
	fileRepo := repo.NewLocalFileRepo(repo.WithLocalFilePath("/tmp"))

	metadataRepo := repo.NewFileMetadataRepo(db)

	return domain.NewFileDO(fileRepo, metadataRepo)
}
