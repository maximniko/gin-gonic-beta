package models

import (
	"database/sql"
	"github.com/doug-martin/goqu/v9"
)

const Instance string = "appInstance"

func NewApp() *App {
	return &App{}
}

type App struct {
	db   *sql.DB
	goqu *goqu.Database
}

func (a *App) SetDbRead(db *sql.DB) {
	a.db = db
}

func (a *App) GetDbRead() *sql.DB {
	return a.db
}
