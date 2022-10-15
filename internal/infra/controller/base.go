package controller

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type BaseHandler struct {
	db *sql.DB
}

func NewBaseHandler(db *sql.DB) *BaseHandler {
	return &BaseHandler{db: db}
}
