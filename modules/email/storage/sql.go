package storage

import "gorm.io/gorm"

type SqlModel struct {
	db *gorm.DB
}

func NewSqlModel(db *gorm.DB) *SqlModel {
	return &SqlModel{db: db}
}
