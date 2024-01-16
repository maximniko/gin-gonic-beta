package db

import "github.com/doug-martin/goqu/v9"

type MysqlParams interface {
	Apply(sd *goqu.SelectDataset) *goqu.SelectDataset
	Append(p string, v interface{}) MysqlParams
}
