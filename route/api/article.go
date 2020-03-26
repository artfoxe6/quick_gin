package api

import (
	"github.com/gin-gonic/gin"
	"quick_gin/route/pagecache"
	"quick_gin/service/ArticleService"
	"quick_gin/util/request"
)

func LoadArticleRoute(r *gin.Engine) {

	r.POST("/articl e/add", func(c *gin.Context) {
		ArticleService.Add(request.New(c))
	})


	r.GET("/article/detail", pagecache.Second(300, func(c *gin.Context) {
		ArticleService.Detail(request.New(c))
	}))

}
