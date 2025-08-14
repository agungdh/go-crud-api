package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProjectHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"surimbim": "dududuw"})
	}
}
