package entity

import "gorm.io/gorm"

type Heroes struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
}
