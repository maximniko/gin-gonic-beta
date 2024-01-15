package params

import (
	"github.com/doug-martin/goqu/v9"
)

type PkwTecdocArticleSrc struct {
	params map[string]interface{}
}

func (r *PkwTecdocArticleSrc) Apply(sd *goqu.SelectDataset) *goqu.SelectDataset {
	if value, exists := r.getParam("articleId"); exists {
		sd.Where(goqu.C("articleId").Eq(value.(int)))
	}

	return sd
}

func (r *PkwTecdocArticleSrc) Append(p string, v interface{}) *PkwTecdocArticleSrc {
	r.params[p] = v
	return r
}

func (r *PkwTecdocArticleSrc) getParam(key string) (interface{}, bool) {
	value, exists := r.params[key]
	return value, exists
}
