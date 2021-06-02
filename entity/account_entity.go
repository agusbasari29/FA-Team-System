package entity

import "gorm.io/gorm"

type Accounts struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
}
