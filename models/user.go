package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	Firstname string
	Lastname  string
	Email     string `gorm:"type:varchar(255);unique_index"`
	Phone     string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t *User) GetCount(db *gorm.DB) int {
	var count int64
	db.Model(&t).Count(&count)
	return int(count)
}

func (t *User) GetEmailCount(db *gorm.DB, email string) int {
	var count int64
	db.Model(&t).Where("email = ?", email).Count(&count)
	return int(count)
}
