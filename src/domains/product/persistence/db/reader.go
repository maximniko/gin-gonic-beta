package db

import (
	repo "awesomeProject/src/domains/product/persistence/db/repositories"
	"awesomeProject/src/domains/product/persistence/db/scheme"
	"awesomeProject/src/interfaces/persistence/repository/db"
	"database/sql"
)

type Reader struct {
	Db *sql.DB
}

func (r *Reader) GetFirst(p db.MysqlParams) *scheme.PkwTecdocArticleSrc {
	repository := &repo.PkwTecdocArticleSrc{Db: r.Db}

	return repository.First(p)
}

func (r *Reader) GetList(p db.MysqlParams) *[]scheme.PkwTecdocArticleSrc {
	repository := &repo.PkwTecdocArticleSrc{Db: r.Db}

	return repository.List(p)
}
