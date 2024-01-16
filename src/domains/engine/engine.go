package engine

import (
	"awesomeProject/src/http/actions"
	"awesomeProject/src/http/middlewares"
	"github.com/gin-gonic/gin"
)

func MakeEngine() (e *gin.Engine) {
	// Disable Console Color
	// gin.DisableConsoleColor()
	e = gin.Default()
	e.Use(middlewares.AppendApp())

	actions.InitRoutes(e)

	return
}
