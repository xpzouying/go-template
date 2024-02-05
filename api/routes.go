package api

import (
	"embed"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/xpzouying/go-template/api/response"
	"github.com/xpzouying/go-template/internal/service"
	"github.com/xpzouying/go-template/pkg/embed_folder"
)

func RegisterService(
	r *gin.Engine,
	service service.Service,
	indexPage []byte,
	buildFS embed.FS,
) {
	setWebRouter(r, buildFS, indexPage)

	r.GET("/status", makeStatus(service))
}

func makeStatus(s service.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		result := s.Status(ctx)

		response.WriteSuccess(c, result)
	}
}

func setWebRouter(r *gin.Engine, buildFS embed.FS, indexPage []byte) {

	r.Use(static.Serve("/", embed_folder.EmbedFolder(buildFS, "web/dist")))

	r.NoRoute(func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html", indexPage)
	})
}
