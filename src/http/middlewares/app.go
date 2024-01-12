package middlewares

import (
	"awesomeProject/src/domains/app/models"
	"github.com/gin-gonic/gin"
)

func AppendApp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(models.Instance, &models.App{})
		c.Next()
	}
}
