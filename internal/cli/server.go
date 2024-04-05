package cli

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xpzouying/go-cmd-project-template/internal/config"
)

// https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
func startHTTPServer(ctx context.Context, cfg *config.Config) error {
	cliLogger.Infof("Starting HTTP server on %s", cfg.ListenAddr)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	svr := &http.Server{
		Addr:    cfg.ListenAddr,
		Handler: r,
	}

	go func() {
		// service connections
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			cliLogger.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	<-ctx.Done()

	cliLogger.Info("Shutting down HTTP server...")

	// NOTE(zy): use a new context to control the timeout.
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := svr.Shutdown(shutdownCtx); err != nil {
		return err
	}

	cliLogger.Info("HTTP server stopped")
	return nil
}
