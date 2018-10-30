package actions

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Root of this project
func Root(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
