package v1

import (
	"awesomeProject/src/http/middlewares/product"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ActionProduct(c *gin.Context) {
	productInstance, _ := c.Get(product.InstanceProduct)

	c.JSON(http.StatusOK, productInstance)
}
