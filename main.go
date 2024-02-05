package main

import (
	"embed"
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xpzouying/go-template/api"
	"github.com/xpzouying/go-template/internal/service"
)

//go:embed web/dist/*
var buildFS embed.FS

//go:embed web/dist/index.html
var indexPage []byte

func main() {
	var (
		port string
	)
	flag.StringVar(&port, "port", ":8080", "port to listen on")
	flag.Parse()

	r := gin.New()
	s := service.New()

	api.RegisterService(r, s, indexPage, buildFS)

	logrus.Infof("listening on port %v", port)
	logrus.Fatal(r.Run(port))
}
