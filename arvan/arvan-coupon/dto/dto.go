package dto

import "arvan-coupon/entity"

type CouponResponse struct {
	Code     string `json:"code"`
	Capacity uint64 `json:"capacity"`
	Value    uint64 `json:"value"`
}

func ToCouponResponse(coupon entity.Coupon) CouponResponse {
	return CouponResponse{
		Code:     coupon.Code,
		Capacity: coupon.Capacity,
		Value:    coupon.Value,
	}
}
