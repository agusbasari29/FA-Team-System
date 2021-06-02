package entity

import "gorm.io/gorm"

type Slp struct {
	gorm.Model
	ID int `json:"id" gorm:"primaryKey;autoIncrement"`
}
