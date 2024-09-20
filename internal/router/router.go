package router

import (
	"github.com/xpzouying/go-cmd-project-template/internal/api"
	"github.com/xpzouying/go-cmd-project-template/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {

	r.Use(middleware.GinZapLogger())

	r.GET("/status", api.HandleGetStatus)

	setAPIRouter(r)
	setUserRouter(r)
}

func setAPIRouter(r *gin.Engine) {
	apiRouter := r.Group("/api")

	apiRouter.GET("/status", api.HandleGetStatus)
}

func setUserRouter(r *gin.Engine) {

	userRouter := r.Group("/user")

	userRouter.POST("/create", api.HandleCreateUser)
	userRouter.GET("/get", api.HandleGetUser)
}
