package models

import "database/sql"

const Instance string = "appInstance"

type App struct {
	db *sql.DB
}

func (a *App) SetDbRead(db *sql.DB) {
	a.db = db
}

func (a *App) GetDbRead() *sql.DB {
	return a.db
}
