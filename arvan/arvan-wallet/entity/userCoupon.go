package entity

import "gorm.io/gorm"

type UserCoupon struct {
	gorm.Model
	ID     uint   `gorm:"primary_key"`
	Code   string `gorm:"column:code"`
	UserID uint
}
