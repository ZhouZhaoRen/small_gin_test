package app

import "github.com/gin-gonic/gin"

type Handler interface {
	Router() string
	Handler(ctx *gin.Context)
}
