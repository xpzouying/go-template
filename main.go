package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xpzouying/go-template/api"
	"github.com/xpzouying/go-template/internal/service"
)

func main() {
	var (
		port string
	)
	flag.StringVar(&port, "port", ":8080", "port to listen on")
	flag.Parse()

	r := gin.New()
	s := service.New()

	api.RegisterService(r, s)

	logrus.Infof("listening on port %v", port)
	logrus.Fatal(r.Run(port))
}
