package api

import (
	"github.com/gin-gonic/gin"
	"quick_gin/route/pagecache"
	"quick_gin/service/article"
	"quick_gin/util/request"
)

func LoadArticleRoute(r *gin.Engine) {

	r.POST("/article/add", func(c *gin.Context) {
		article.Add(request.New(c))
	})

	r.GET("/article/detail", pagecache.Second(300, func(c *gin.Context) {
		article.Detail(request.New(c))
	}))

}
