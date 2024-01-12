package v1

import (
	"github.com/gin-gonic/gin"
)

func ActionPing(c *gin.Context) {
	c.String(200, "pong")
}
