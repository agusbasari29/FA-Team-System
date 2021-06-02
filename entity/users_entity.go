package entity

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
}
