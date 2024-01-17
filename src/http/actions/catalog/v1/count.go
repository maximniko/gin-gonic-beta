package v1

import (
	"awesomeProject/src/domains/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ActionCatalogCount(c *gin.Context) {

	var count int
	err := c.MustGet(models.Instance).(*models.App).GetDbRead().QueryRow("SELECT COUNT(*) FROM pkw_frontend.mix_catalog_tree where skinId = 2 and type = 3").Scan(&count)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"albums": count})
}

type CatalogCount struct {
	count int
}
