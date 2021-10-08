package sqlLiteRepo

import (
	"arvan-coupon/dal/connection"
	"arvan-coupon/entity"
	"errors"
)

type CouponRepository struct {
	DB connection.Database
}

func (u CouponRepository) SaveCoupon(coupon entity.Coupon) error {
	u.DB.Create(&coupon)
	return nil
}

func (u CouponRepository) GetCouponByCode(code string) (entity.Coupon, error) {
	var coupon entity.Coupon
	result := u.DB.Model(&entity.Coupon{}).Where("code = ?", code).Scan(&coupon)
	if result.RowsAffected == 0 {
		return coupon, errors.New("no coupon found")
	}
	return coupon, nil
}

func NewCouponRepository() CouponRepository {
	return CouponRepository{
		DB: connection.GetInstance(connection.GetSqlLiteConfig()),
	}
}
