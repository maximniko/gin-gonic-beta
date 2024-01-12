package actions

import (
	"awesomeProject/src/http/actions/ping"
	"awesomeProject/src/http/actions/product"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	initRoutesV1(e)
}

func initRoutesV1(e *gin.Engine) {
	v1 := e.Group("/v1")
	{
		ping.RoutesPingV1(v1)
		product.RoutesProductV1(v1)
	}
}
