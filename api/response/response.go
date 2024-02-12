package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"result"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data"`
}

// WriteSuccess writes a success response.
func WriteSuccess(c *gin.Context, data interface{}) {

	c.JSON(http.StatusOK,
		Response{
			Success: true,
			Data:    data,
		})
}

func WriteError(c *gin.Context, errMsg string) {
	c.JSON(http.StatusOK,
		Response{
			Success: false,
			Error:   errMsg,
		})
}

type StatusResult struct {
	Version   string `json:"version"`
	StartTime string `json:"start_time"`
}
