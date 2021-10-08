package sqlLiteRepo

import (
	"arvan-wallet/dal/connection"
	"arvan-wallet/entity"
	"errors"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	DB connection.Database
}

func (u UserRepository) SaveUser(user entity.User) error {
	result := u.DB.Create(&user)
	return result.Error
}

func (u UserRepository) UpdateUser(usr entity.User) error {
	var user entity.User
	if result := u.DB.Find(&user, "phone_number = ?", usr.PhoneNumber); result.Error != nil {
		return result.Error
	}
	user.Balance = usr.Balance
	result := u.DB.Save(&user)
	return result.Error
}

func (u UserRepository) GetUsers() ([]entity.User, error) {
	var users []entity.User
	result := u.DB.Find(&users)
	return users, result.Error
}

func (u UserRepository) GetUserByPhoneNumber(phoneNumber string) (entity.User, error) {
	var user entity.User
	result := u.DB.Find(&user, "phone_number = ?", phoneNumber)
	if result.RowsAffected == 0 {
		return entity.User{}, errors.New("user not found")
	}
	return user, result.Error
}

func (u UserRepository) GetUsersByCoupon(coupon entity.UserCoupon) ([]entity.User, error) {
	var users []entity.User
	result := u.DB.Model(&entity.User{}).Joins("join user_coupons on user_coupons.code = ? and user_coupons.user_id = users.id", coupon.Code).Scan(&users)
	return users, result.Error
}

func (u UserRepository) CountUsersByCoupon(coupon entity.UserCoupon) (uint16, error) {
	var count int64
	result := u.DB.Model(&entity.User{}).Joins("join user_coupons on user_coupons.code = ? and user_coupons.user_id = users.id", coupon.Code).Count(&count)
	return uint16(count), result.Error
}

func NewUserRepository() UserRepository {
	return UserRepository{
		DB: connection.GetInstance(connection.GetSqlLiteConfig()),
	}
}

func NewTxnUserRepository(db *gorm.DB) UserRepository {
	if db == nil {
		log.Print("Transaction Database not found")
		return UserRepository{}
	}
	return UserRepository{
		DB: connection.Database{
			DB: db,
		},
	}
}
