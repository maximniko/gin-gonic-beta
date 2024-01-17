package components

import (
	"awesomeProject/src/domains/app/models"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"log"
)

func NewGoqu(app *models.App) *goqu.Database {
	qodu := goqu.New("mysql8", app.GetDbRead())
	qodu.Logger(log.Default())
	fmt.Println("init NewGoqu() gin.HandlerFunc")
	return qodu
}
