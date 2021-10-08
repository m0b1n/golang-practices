package entity

import "gorm.io/gorm"

type Coupon struct {
	gorm.Model
	ID       uint   `gorm:"primary_key"`
	Code     string `gorm:"column:code"`
	Capacity uint64 `gorm:"column:capacity"`
	Value    uint64 `gorm:"column:calue"`
}
