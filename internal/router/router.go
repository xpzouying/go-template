package router

import (
	"github.com/xpzouying/go-cmd-project-template/internal/controller"
	"github.com/xpzouying/go-cmd-project-template/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {

	r.Use(middleware.GinZapLogger())

	r.GET("/status", controller.HandleGetStatus)

	setAPIRouter(r)
}

func setAPIRouter(r *gin.Engine) {
	apiRouter := r.Group("/api")

	apiRouter.GET("/status", controller.HandleGetStatus)
}
