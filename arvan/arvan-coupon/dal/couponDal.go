package dal

import "arvan-coupon/entity"

type UserCouponRepositoryInterface interface {
	SaveCoupon(coupon entity.Coupon) error
	GetCouponByCode(code string) (entity.Coupon, error)
}
