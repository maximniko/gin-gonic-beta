package engine

import (
	"awesomeProject/src/domains/app/models"
	"awesomeProject/src/http/actions"
	"awesomeProject/src/http/middlewares"
	"github.com/gin-gonic/gin"
)

func MakeEngine(app *models.App) (e *gin.Engine) {
	// Disable Console Color
	// gin.DisableConsoleColor()
	e = gin.Default()
	e.Use(middlewares.AppendApp(app))

	actions.InitRoutes(e)

	return
}
