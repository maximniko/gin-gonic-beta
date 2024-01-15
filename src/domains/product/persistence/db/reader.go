package db

import (
	repo "awesomeProject/src/domains/product/persistence/db/repositories"
	"awesomeProject/src/domains/product/persistence/db/repositories/params"
	"awesomeProject/src/domains/product/persistence/db/scheme"
	"database/sql"
)

type Reader struct {
	Db *sql.DB
}

func (r *Reader) GetFirst(p *params.PkwTecdocArticleSrc) *scheme.PkwTecdocArticleSrc {
	repository := &repo.PkwTecdocArticleSrc{Db: r.Db}

	return repository.First(p)
}
