package entity

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID          uint         `gorm:"primary_key"`
	PhoneNumber string       `gorm:"column:phone_number;not null"`
	Balance     uint64       `gorm:"column:balance"`
	UserCoupons []UserCoupon `gorm:"ForeignKey:UserID"`
	UserTxns    []UserTxn    `gorm:"ForeignKey:UserID"`
}

func (user User) String() string {
	return fmt.Sprintf("User [ID: %d, PhoneNumber: %s, Balance: %d]", user.ID, user.PhoneNumber, user.Balance)
}
