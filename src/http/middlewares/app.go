package middlewares

import (
	"awesomeProject/src/domains/app/models"
	"github.com/gin-gonic/gin"
)

func AppendApp(app *models.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(models.Instance, app)
		c.Next()
	}
}
