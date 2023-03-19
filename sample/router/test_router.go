package router

import (
	"github.com/gin-gonic/gin"
)

func TestRouter(g *gin.Engine) {
	g.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})
}
