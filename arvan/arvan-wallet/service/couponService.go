package service

import (
	"arvan-wallet/dal"
	"arvan-wallet/dal/postgresRepo"
	"arvan-wallet/entity"
	"gorm.io/gorm"
)

type CouponService struct {
	userDal dal.UserRepositoryInterface
}

type CouponServiceInterface interface {
	FetchCouponUsersDetail(coupon string) ([]entity.User, error)
}

func (service CouponService) FetchCouponUsersDetail(coupon string) ([]entity.User, error) {
	return service.userDal.GetUsersByCoupon(entity.UserCoupon{Code: coupon})
}

func NewCouponServiceInstance() CouponService {
	return CouponService{
		userDal: postgresRepo.NewUserRepository(),
	}
}

func NewTxnCouponServiceInstance(db *gorm.DB) CouponService {
	return CouponService{
		userDal: postgresRepo.NewTxnUserRepository(db),
	}
}
