package restController

import (
	"arvan-wallet/common"
	"arvan-wallet/dto"
	"arvan-wallet/service"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UsersRouter(router *gin.RouterGroup) {
	router.GET("/:phonenumber", GetUserInfo())
	router.POST("/submitcoupon", common.DBTransactionMiddleware(), SubmitUserCoupon())
	router.POST("/submituser", SubmitUser())
}

func GetUserInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		phoneNumber, flg := c.Params.Get("phonenumber")
		if !flg {
			c.JSON(http.StatusBadRequest, gin.H{"error": "not enough parameters"})
			return
		}
		user, err := service.NewUserServiceInstance().GetUserDetail(phoneNumber)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, dto.ToUserResponse(user))

	}
}

func SubmitUserCoupon() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqBody dto.CouponSubmitRequest
		if err := c.ShouldBind(&reqBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		db := c.MustGet("db_trx").(*gorm.DB)

		err := service.NewTxnUserServiceInstance(db).InsertCoupon(reqBody.PhoneNumber, reqBody.Code)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "service is not available"})
			//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, reqBody)
	}
}

func SubmitUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		res := dto.UserResponse{}
		if err := c.ShouldBind(&res); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := service.NewUserServiceInstance().InsertUser(res); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"msg": "user created successfully"})
	}
}
