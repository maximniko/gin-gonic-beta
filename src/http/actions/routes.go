package actions

import (
	"awesomeProject/src/http/actions/catalog"
	"awesomeProject/src/http/actions/ping"
	"awesomeProject/src/http/actions/product"
	"github.com/gin-gonic/gin"
)

func InitRoutes(e *gin.Engine) {
	initRoutesV1(e)
}

func initRoutesV1(e *gin.Engine) {
	ping.RoutesPingV1(e)
	v1 := e.Group("/v1")
	{
		product.RoutesProductV1(v1)
		catalog.RoutesCatalogV1(v1)
	}
}
