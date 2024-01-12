package db

import (
	"awesomeProject/src/domains/product/models"
	"database/sql"
)

type Reader struct {
	Db *sql.DB
}

func (r *Reader) GetById(id int) *models.Product {
	//make request to db and translate it to models.Product
	return &models.Product{Id: id}
}
