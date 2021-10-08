package dal

import "arvan-wallet/entity"

type UserCouponRepositoryInterface interface {
	SaveCoupon(coupon entity.UserCoupon) error
	GetCouponsByUser(user entity.User) ([]entity.UserCoupon, error)
	GetCouponsCount(coupon string) (int64, error)
}
