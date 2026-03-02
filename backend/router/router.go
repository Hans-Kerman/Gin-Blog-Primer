package router

import (
	"github.com/Hans-Kerman/GinBlogPrimer/backend/controllers"
	"github.com/Hans-Kerman/GinBlogPrimer/backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	auth := r.Group("/api/auth")
	{
		auth.POST("/login", controllers.Login)
		auth.POST("/register", controllers.Register)
	}

	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRates) //不需要登录注册就可以使用的功能
	api.Use(middlewares.AuthMiddleWare())                   //大括号内所有路由需要经过中间件
	{
		api.POST("/exchangeRates", controllers.CreateExchangeRate)
	}
	return r
}
