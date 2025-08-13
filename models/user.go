package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"column:name;size:255"`
	Email    string `gorm:"column:email;not null"`
	Password string `gorm:"column:password;not null"`
}