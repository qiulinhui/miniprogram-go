package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint
	Name      string
	Author    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (book *Book) Insert() error {
	return DB.Create(&book).Error
}

func (book *Book) SelectByID() error {
	return DB.First(&book, book.ID).Error
}
