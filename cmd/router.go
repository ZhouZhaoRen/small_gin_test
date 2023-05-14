package main

import (
	"github.com/gin-gonic/gin"
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

	apiConcern := r.Group("/api/concern")
	//apiv1.Use(jwt.JWT()) // 在这里添加中间件，每次执行都会先走这里进行token验证
	{
		// 对标签的操作
		// 获取标签列表
		apiConcern.GET("/tags")
	}

	return r
}
