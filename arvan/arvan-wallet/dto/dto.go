package dto

import "arvan-wallet/entity"

type UserResponse struct {
	PhoneNumber string `json:"phoneNumber"`
	Balance     uint64 `json:"balance"`
}

type UserResponse1 struct {
	PhoneNumber string `json:"phoneNumber"`
}

type CouponResponse struct {
	Code     string `json:"code"`
	Capacity uint64 `json:"capacity"`
	Value    uint64 `json:"value"`
}

func ToUserResponse(user entity.User) UserResponse {
	return UserResponse{
		PhoneNumber: user.PhoneNumber,
		Balance:     user.Balance,
	}
}

func ToUserResponse1(user entity.User) UserResponse1 {
	return UserResponse1{
		PhoneNumber: user.PhoneNumber,
	}
}

func ToUserResponses(users []entity.User) []UserResponse1 {
	var userResponses []UserResponse1
	for _, user := range users {
		userRes := ToUserResponse1(user)
		userResponses = append(userResponses, userRes)
	}
	return userResponses
}

type CouponSubmitRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Code        string `json:"code"`
}
