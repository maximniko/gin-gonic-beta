package product

import (
	"awesomeProject/src/domains/app/models"
	"awesomeProject/src/domains/product/persistence/db"
	"awesomeProject/src/domains/product/persistence/db/repositories/params"
	"awesomeProject/src/http/handlers"
	"awesomeProject/src/http/handlers/panics"
	"github.com/gin-gonic/gin"
	"strconv"
)

const ParamArticleId = "articleId"
const InstanceProduct = "instanceArticle"

func ArticleIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := getIdFromRequest(c)
		trySetInstanceProduct(c, id)

		c.Next()
	}
}

func getIdFromRequest(c *gin.Context) int {
	defer handlers.HandleRuntimePanic(c)

	articleId, ok := c.Params.Get(ParamArticleId)

	if !ok {
		panic(panics.NewBadRequest())
	}

	id, err := strconv.Atoi(articleId)

	if err != nil || id < 1 {
		panic(panics.NewBadRequest())
	}

	return id
}

func trySetInstanceProduct(c *gin.Context, id int) {
	defer handlers.HandleRuntimePanic(c)

	reader := db.Reader{Db: c.MustGet(models.Instance).(*models.App).GetDbRead()}

	p := params.NewPkwTecdocArticleSrc()

	c.Set(InstanceProduct, reader.GetFirst(p.Append("articleId", id)))
}
