package main

import (
	"github.com/Hans-Kerman/GinBlogPrimer/backend/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitConfig()
	r := gin.Default()
	r.GET("ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(config.AppConfig.App.Port)
}
