package entity

import "gorm.io/gorm"

type Withdraw struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
}
