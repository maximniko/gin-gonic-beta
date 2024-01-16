package repositories

import (
	"awesomeProject/src/domains/product/persistence/db/scheme"
	"awesomeProject/src/http/handlers/panics"
	"awesomeProject/src/interfaces/persistence/repository/db"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"log"
)

type PkwTecdocArticleSrc struct {
	Db *sql.DB
}

func (r *PkwTecdocArticleSrc) First(p db.MysqlParams) *scheme.PkwTecdocArticleSrc {
	var product scheme.PkwTecdocArticleSrc

	found, err := p.Apply(r.selectDataset()).
		Limit(1).
		ScanStruct(&product)

	if err != nil {
		log.Println(err.Error())
		p := panics.NewBadRequest()
		panic(p)
	}

	if !found {
		log.Println("Product not found")
		p := panics.NewBadRequest()
		panic(p)
	}

	return &product
}

func (r *PkwTecdocArticleSrc) List(p db.MysqlParams) *[]scheme.PkwTecdocArticleSrc {
	var products []scheme.PkwTecdocArticleSrc

	err := p.Apply(r.selectDataset()).ScanStructs(&products)

	if err != nil {
		panic(err.Error())
	}

	return &products
}

func (r *PkwTecdocArticleSrc) selectDataset() *goqu.SelectDataset {
	q := goqu.New("mysql8", r.Db)
	t := goqu.S("pkwteile_store").Table("pkw_tecdoc_article_src")

	return q.From(t)
}
