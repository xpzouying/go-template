package api

import (
	"github.com/gin-gonic/gin"
	"github.com/xpzouying/go-template/api/response"
	"github.com/xpzouying/go-template/internal/service"
)

func RegisterService(r gin.IRouter, service service.Service) {

	r.GET("/status", makeStatus(service))
}

func makeStatus(s service.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		result := s.Status(ctx)

		response.WriteSuccess(c, result)
	}
}
