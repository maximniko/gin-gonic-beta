package app

import (
	"awesomeProject/src/domains/app/models"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func MysqlConnectRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := os.Getenv("MYSQL_USERNAME")
		password := os.Getenv("MYSQL_PASSWORD")
		dbName := os.Getenv("MYSQL_DATABASE")

		dataSourceName := fmt.Sprintf("%s:%s@/%s", username, password, dbName)

		fmt.Println(dataSourceName)
		// Открываем соединение с базой данных
		db, err := sql.Open("mysql", dataSourceName)
		if err != nil {
			panic(err.Error())
		}

		err = db.Ping()
		if err != nil {
			panic(err.Error())
		}

		c.MustGet(models.Instance).(*models.App).SetDbRead(db)
		c.Next()
	}
}
