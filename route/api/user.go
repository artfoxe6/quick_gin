package api

import (
	"github.com/gin-gonic/gin"
	"quick_gin/middleware"
	"quick_gin/service/user"
	"quick_gin/util/request"
)

func LoadUserRoute(r *gin.Engine) {
	g := r.Group("/user")

	//添加用户
	g.POST("/add", func(context *gin.Context) {
		user.Add(request.New(context))
	})

	//用户列表
	g.GET("/list", func(context *gin.Context) {
		user.List(request.New(context))
	})

	//用户列表以及用户的文章
	g.GET("/list_with_article", func(c *gin.Context) {
		user.ListWithArticles(request.New(c))
	})

	//获取token
	g.GET("/token", func(context *gin.Context) {
		user.CreateToken(request.New(context))
	})

	//需要token认证才能请求的路由
	auth := r.Group("/user/info")
	auth.Use(middleware.Auth())
	{
		auth.GET("/", func(context *gin.Context) {
			user.Info(request.New(context))
		})
	}

}
