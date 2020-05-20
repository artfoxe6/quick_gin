package api

import (
	"github.com/gin-gonic/gin"
	"quick_gin/route/pagecache"
	"quick_gin/service/article"
	"quick_gin/util/request"
)

func LoadArticleRoute(r *gin.Engine) {

	//添加文章
	r.POST("/article/add", func(c *gin.Context) {
		article.Add(request.New(c))
	})

	//获取文章详情
	//此接口将缓存300秒
	r.GET("/article/detail", pagecache.Second(300, func(c *gin.Context) {
		article.Detail(request.New(c))
	}))

}
