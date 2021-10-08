package postgresRepo

import (
	"arvan-wallet/dal/connection"
	"arvan-wallet/entity"
	"gorm.io/gorm"
	"log"
)

type UserCouponRepository struct {
	DB connection.Database
}

func (u UserCouponRepository) SaveCoupon(coupon entity.UserCoupon) error {
	result := u.DB.Create(&coupon)
	return result.Error
}

func (u UserCouponRepository) GetCouponsByUser(user entity.User) ([]entity.UserCoupon, error) {
	var coupons []entity.UserCoupon
	result := u.DB.Model(&entity.UserCoupon{}).Joins("join users on users.phone_number = ?", user.PhoneNumber).Scan(&coupons)
	return coupons, result.Error
}

func (u UserCouponRepository) GetCouponsCount(coupon string) (int64, error) {
	var count int64
	result := u.DB.Model(&entity.UserCoupon{}).Where("code = ?", coupon).Count(&count)
	return count, result.Error
}

func NewUserCouponRepository() UserCouponRepository {
	return UserCouponRepository{
		DB: connection.GetPostgresInstance(connection.GetPostgresConfig()),
	}
}

func NewTxnUserCouponRepository(db *gorm.DB) UserCouponRepository {
	if db == nil {
		log.Print("Transaction Database not found")
		return UserCouponRepository{}
	}
	return UserCouponRepository{
		DB: connection.Database{
			DB: db,
		},
	}
}
