package restController

import (
	"arvan-wallet/dto"
	"arvan-wallet/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CouponRouter(router *gin.RouterGroup) {
	router.GET("/:coupon", GetCouponDetail())
}

func GetCouponDetail() gin.HandlerFunc {
	return func(c *gin.Context) {
		q, _ := c.Params.Get("coupon")
		users, err := service.NewCouponServiceInstance().FetchCouponUsersDetail(q)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if len(users) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no users"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"amount":      len(users),
			"listNumbers": dto.ToUserResponses(users),
		})
	}
}
