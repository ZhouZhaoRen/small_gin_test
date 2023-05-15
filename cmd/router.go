package main

import (
	"github.com/gin-gonic/gin"
	"small_gin_test/app"
	"small_gin_test/app/uid_flake"
	//v1 "go-gin-example/routers/api/v1"
)

// InitRouter 初始化
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	hanlders := []app.Handler{
		&uid_flake.UidFlake{},
	}
	for i := range hanlders {
		r.POST(hanlders[i].Router(), hanlders[i].Handler)
	}

	apiConcern := r.Group("/api/concern")
	//apiv1.Use(jwt.JWT()) // 在这里添加中间件，每次执行都会先走这里进行token验证
	{
		// 对标签的操作
		// 获取标签列表
		apiConcern.GET("/tags")
	}

	return r
}
