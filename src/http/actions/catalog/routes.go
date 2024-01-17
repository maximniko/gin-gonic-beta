package catalog

import (
	v1 "awesomeProject/src/http/actions/catalog/v1"
	"github.com/gin-gonic/gin"
)

func RoutesCatalogV1(g *gin.RouterGroup) {
	// Product test
	g.GET("/catalog/count", v1.ActionCatalogCount)
}
