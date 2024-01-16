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
		dns := dns()

		fmt.Println(dns)
		// Открываем соединение с базой данных
		db, err := sql.Open("mysql", dns)
		if err != nil {
			panic(err.Error())
		}

		c.MustGet(models.Instance).(*models.App).SetDbRead(db)
		c.Next()
	}
}

func dns() string {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	hostname := os.Getenv("MYSQL_HOSTNAME")

	return fmt.Sprintf("%s:%s@tcp(%s)/", username, password, hostname)
}
