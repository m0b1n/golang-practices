package service

import (
	"arvan-wallet/dal"
	"arvan-wallet/dal/postgresRepo"
	"arvan-wallet/dto"
	"arvan-wallet/entity"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"gorm.io/gorm"
)

type UserService struct {
	userDal   dal.UserRepositoryInterface
	couponDal dal.UserCouponRepositoryInterface
	txnDal    dal.UserTxnRepositoryInterface
}

type UserServiceInterface interface {
	InsertCoupon(userPhoneNumber string, coupon string) error
	InsertUser(user dto.UserResponse) error
	GetUserDetail(userPhoneNumber string) (entity.User, error)
}

func (service UserService) GetUserDetail(userPhoneNumber string) (entity.User, error) {
	if user, err := service.userDal.GetUserByPhoneNumber(userPhoneNumber); err == nil {
		return user, nil
	} else {
		return entity.User{}, err
	}
}

func (service UserService) InsertUser(user dto.UserResponse) error {
	return service.userDal.SaveUser(entity.User{
		PhoneNumber: user.PhoneNumber,
		Balance:     user.Balance,
	})
}

func (service UserService) InsertCoupon(userPhoneNumber string, coupon string) error {
	// first check coupon validity
	// we should ask coupon service
	// what happens if this trx fails
	response, err := http.Get(fmt.Sprintf("http://coupon:8080/coupon/%s", coupon))
	if err != nil || response.StatusCode != 200 {
		return errors.New("not valid coupon")
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	var couponRes dto.CouponResponse
	if err := json.Unmarshal(responseData, &couponRes); err != nil {
		return err
	}
	// Should search for users with related coupon and match it with cap of coupon service
	count, err := service.couponDal.GetCouponsCount(coupon)
	if err != nil {
		return err
	}
	if count >= int64(couponRes.Capacity) {
		return errors.New("limit exceeded")
	}
	user, err := service.userDal.GetUserByPhoneNumber(userPhoneNumber)
	if err != nil {
		return err
	}
	// then should save coupon for user
	if err := service.couponDal.SaveCoupon(entity.UserCoupon{
		Code:   coupon,
		UserID: user.ID,
	}); err != nil {
		return err
	}
	//then should add user transaction
	if err := service.txnDal.SaveUserTxn(entity.UserTxn{
		Amount: couponRes.Value,
		UserID: user.ID,
	}); err != nil {
		return err
	}
	// then update user balance
	if err := service.userDal.UpdateUser(entity.User{
		PhoneNumber: userPhoneNumber,
		Balance:     couponRes.Value + user.Balance,
	}); err != nil {
		return err
	}
	return nil
}

func NewUserServiceInstance() UserService {
	return UserService{
		userDal:   postgresRepo.NewUserRepository(),
		couponDal: postgresRepo.NewUserCouponRepository(),
		txnDal:    postgresRepo.NewUserTxnRepository(),
	}
}

func NewTxnUserServiceInstance(db *gorm.DB) UserService {
	return UserService{
		userDal:   postgresRepo.NewTxnUserRepository(db),
		couponDal: postgresRepo.NewTxnUserCouponRepository(db),
		txnDal:    postgresRepo.NewTxnUserTxnRepository(db),
	}
}
