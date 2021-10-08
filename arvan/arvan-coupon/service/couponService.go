package service

import (
	"arvan-coupon/dal"
	"arvan-coupon/dal/sqlLiteRepo"
	"arvan-coupon/dto"
	"arvan-coupon/entity"
)

type CouponService struct {
	couponDal dal.UserCouponRepositoryInterface
}

type CouponServiceInterface interface {
	FetchCouponDetails(code string) (entity.Coupon, error)
	SubmitNewCoupon(coupon dto.CouponResponse) error
}

func (service CouponService) FetchCouponDetails(code string) (entity.Coupon, error) {
	return service.couponDal.GetCouponByCode(code)
}

func (service CouponService) SubmitNewCoupon(coupon dto.CouponResponse) error {
	return service.couponDal.SaveCoupon(entity.Coupon{
		Code:     coupon.Code,
		Capacity: coupon.Capacity,
		Value:    coupon.Value,
	})
}

func NewCouponServiceInstance() CouponService {
	return CouponService{
		couponDal: sqlLiteRepo.NewCouponRepository(),
	}
}
