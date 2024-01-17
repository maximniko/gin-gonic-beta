package connections

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func NewMysqlConnectRead() *sql.DB {
	dns := dns()

	fmt.Println(dns)
	// Открываем соединение с базой данных
	db, err := sql.Open("mysql", dns)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("init MysqlConnectRead() gin.HandlerFunc")
	return db
}

func dns() string {
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	hostname := os.Getenv("MYSQL_HOSTNAME")

	return fmt.Sprintf("%s:%s@tcp(%s)/", username, password, hostname)
}
