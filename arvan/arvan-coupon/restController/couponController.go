package restController

import (
	"arvan-coupon/dto"
	"arvan-coupon/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CouponRouter(router *gin.RouterGroup) {
	router.GET("/:code", GetCouponDetail())
	router.POST("/", SubmitNewCoupon())
}

func GetCouponDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		q, _ := c.Params.Get("code")
		coupon, err := service.NewCouponServiceInstance().FetchCouponDetails(q)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, dto.ToCouponResponse(coupon))
	}
}

func SubmitNewCoupon() gin.HandlerFunc {
	return func(c *gin.Context) {
		var coupon dto.CouponResponse
		if err := c.ShouldBind(&coupon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := service.NewCouponServiceInstance().SubmitNewCoupon(coupon); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"msg": "coupon created successfully"})
	}
}
