package product

import (
	"awesomeProject/src/http/actions/product/v1"
	"awesomeProject/src/http/middlewares/product"
	"github.com/gin-gonic/gin"
)

func RoutesProductV1(g *gin.RouterGroup) {
	// Product test
	g.GET("/product/:"+product.ParamArticleId, product.ArticleIdMiddleware(), v1.ActionProduct)
}
