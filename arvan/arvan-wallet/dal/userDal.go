package dal

import "arvan-wallet/entity"

type UserRepositoryInterface interface {
	SaveUser(user entity.User) error
	UpdateUser(user entity.User) error
	GetUsers() ([]entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
	GetUsersByCoupon(coupon entity.UserCoupon) ([]entity.User, error)
	CountUsersByCoupon(coupon entity.UserCoupon) (uint16, error)
}
