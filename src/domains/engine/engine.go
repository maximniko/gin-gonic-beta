package engine

import (
	"awesomeProject/src/http/actions"
	"awesomeProject/src/http/middlewares"
	"awesomeProject/src/http/middlewares/app"
	"github.com/gin-gonic/gin"
)

func MakeEngine() (e *gin.Engine) {
	// Disable Console Color
	// gin.DisableConsoleColor()
	e = gin.Default()
	e.Use(middlewares.AppendApp())
	e.Use(app.MysqlConnectRead())

	actions.InitRoutes(e)

	return
}
