package factories

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

func NewGoqu(db *sql.DB) *goqu.Database {
	qodu := goqu.New("mysql8", db)
	//qodu.Logger(log.Default())
	return qodu
}
