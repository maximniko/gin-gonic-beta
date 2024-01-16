package params

import (
	"awesomeProject/src/interfaces/persistence/repository/db"
	"github.com/doug-martin/goqu/v9"
)

func NewPkwTecdocArticleSrc() db.MysqlParams {
	return &PkwTecdocArticleSrc{
		params: make(map[string]interface{}),
	}
}

type PkwTecdocArticleSrc struct {
	params map[string]interface{}
}

func (r *PkwTecdocArticleSrc) Apply(sd *goqu.SelectDataset) *goqu.SelectDataset {
	if value, exists := r.getParam("articleId"); exists {
		sd = sd.Where(goqu.C("articleId").Eq(value.(int)))
	}

	return sd
}

func (r *PkwTecdocArticleSrc) Append(p string, v interface{}) db.MysqlParams {
	r.params[p] = v
	return r
}

func (r *PkwTecdocArticleSrc) getParam(key string) (interface{}, bool) {
	value, exists := r.params[key]
	return value, exists
}
