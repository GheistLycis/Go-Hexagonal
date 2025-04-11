package postgres

import (
	"errors"

	"gorm.io/gorm"
)

type GormAdapter struct {
	db *gorm.DB
}

func NewGormAdapter(db *gorm.DB) *GormAdapter {
	return &GormAdapter{db}
}

func (a *GormAdapter) Get(d any, c ...any) error {
	return a.db.First(d, c...).Error
}

func (a *GormAdapter) Insert(v any) error {
	return a.db.Create(v).Error
}

func (a *GormAdapter) List(d any, c ...any) error {
	return a.db.Find(d, c...).Error
}

func (a *GormAdapter) Update(v any) error {
	result := a.db.Updates(v)

	if result.RowsAffected == 0 {
		return errors.New("entry not found")
	}

	return result.Error
}

func (a *GormAdapter) Upsert(v any) error {
	result := a.db.Updates(v)

	if result.RowsAffected == 0 {
		return a.db.Create(v).Error
	}

	return result.Error
}

func (a *GormAdapter) Query(d any, q string, v ...any) error {
	if d != nil {
		return a.db.Raw(q, v...).Scan(d).Error
	}

	return a.db.Exec(q, v...).Error
}
