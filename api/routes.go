package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	static "github.com/soulteary/gin-static"
	"github.com/xpzouying/go-template/api/response"
	"github.com/xpzouying/go-template/internal/service"
)

func RegisterService(
	r *gin.Engine,
	service *service.Service,
	staticConfig *StaticFSConfig,
) {
	setWebRouter(r, staticConfig)

	setupFileRouter(r, service)

	r.GET("/status", makeStatus(service))

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
}

func makeStatus(s *service.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		result := s.Status(ctx)

		response.WriteSuccess(c, result)
	}
}

func setWebRouter(r *gin.Engine, staticConfig *StaticFSConfig) {

	if staticConfig.Mode.IsDev() {
		r.Use(static.Serve("/", static.LocalFile(staticConfig.StaticFilesPath, false)))
	} else {
		r.NoRoute(
			func(c *gin.Context) {
				if c.Request.URL.Path == "/about" {
					c.Data(
						http.StatusOK,
						"text/html; charset=utf-8",
						[]byte("About Page"),
					)
				}
			},
			static.ServeEmbed("/", staticConfig.embedFS),
		)
	}

	// r.Use(static.Serve("/", embed_folder.EmbedFolder(buildFS, "web/dist")))
}

func setupFileRouter(r *gin.Engine, service *service.Service) {
	subGroup := r.Group("/files")

	subGroup.POST("/upload", service.Upload)
}
