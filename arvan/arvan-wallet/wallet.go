package main

import (
	"arvan-wallet/restController"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "wallet is functioning...")
	})

	v1 := router.Group("/coupon")
	{
		restController.CouponRouter(v1)
	}

	v2 := router.Group("/user")
	{
		restController.UsersRouter(v2)
	}
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
