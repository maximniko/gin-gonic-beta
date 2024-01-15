package repositories

import (
	"awesomeProject/src/domains/product/persistence/db/repositories/params"
	"awesomeProject/src/domains/product/persistence/db/scheme"
	"database/sql"
	"github.com/doug-martin/goqu/v9"
)

type PkwTecdocArticleSrc struct {
	Db *sql.DB
}

func (r *PkwTecdocArticleSrc) First(p *params.PkwTecdocArticleSrc) *scheme.PkwTecdocArticleSrc {
	var product scheme.PkwTecdocArticleSrc

	found, err := p.Apply(r.selectDataset()).
		Limit(1).
		ScanStruct(&product)

	if err != nil {
		panic(err.Error())
	}

	if !found {
		panic("Product not found")
	}

	return &product
}

func (r *PkwTecdocArticleSrc) List(p *params.PkwTecdocArticleSrc) *[]scheme.PkwTecdocArticleSrc {
	var products []scheme.PkwTecdocArticleSrc

	err := p.Apply(r.selectDataset()).ScanStructs(&products)

	if err != nil {
		panic(err.Error())
	}

	return &products
}

func (r *PkwTecdocArticleSrc) selectDataset() *goqu.SelectDataset {
	db := goqu.New("mysql8", r.Db)
	t := goqu.S("pkwteile_store").Table("pkw_tecdoc_article_src")

	return db.From(t)
}
