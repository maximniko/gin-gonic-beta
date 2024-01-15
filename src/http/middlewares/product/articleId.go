package product

import (
	"awesomeProject/src/domains/app/models"
	"awesomeProject/src/domains/product/persistence/db"
	"awesomeProject/src/domains/product/persistence/db/repositories/params"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const ParamArticleId = "articleId"
const InstanceProduct = "instanceArticle"

func ArticleIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		articleId, ok := c.Params.Get(ParamArticleId)

		if !ok {
			c.String(http.StatusBadRequest, "empty product")
			c.Abort()
			return
		}

		id, err := strconv.Atoi(articleId)

		if err != nil || id < 1 {
			c.String(http.StatusBadRequest, "bad product "+articleId)
			c.Abort()
			return
		}

		reader := db.Reader{Db: c.MustGet(models.Instance).(*models.App).GetDbRead()}

		p := params.NewPkwTecdocArticleSrc()
		c.Set(InstanceProduct, reader.GetFirst(p.Append("articleId", id)))
		// Set example variable
		c.Next()
	}
}
