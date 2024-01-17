package ping

import (
	v1 "awesomeProject/src/http/actions/ping/v1"
	"github.com/gin-gonic/gin"
)

func RoutesPingV1(g *gin.RouterGroup) {
	g.GET("/ping", v1.ActionPing)
}
