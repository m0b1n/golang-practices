package main

import (
	"arvan-coupon/restController"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "coupon is functioning...")
	})

	v1 := router.Group("/coupon")
	{
		restController.CouponRouter(v1)
	}

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
