package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name   string `gorm:"not null" json:"name"`
	RollNo int    `gorm:"unique;not null" json:"rollNo"`
	ID     int    `gorm:"primaryKey" json:"id"`
}
