package entity

import "gorm.io/gorm"

type UserTxn struct {
	gorm.Model
	ID     uint   `gorm:"primary_key"`
	Amount uint64 `gorm:"column:amount"`
	UserID uint
}
