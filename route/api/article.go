package api

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	redis "quick_gin/config/cache"
	"quick_gin/service/ArticleService"
	"quick_gin/util/request"
	"time"
)

func LoadArticleRoute(r *gin.Engine) {

	r.POST("/articl e/add", func(c *gin.Context) {
		ArticleService.Add(request.New(c))
	})

	store := persistence.NewRedisCacheWithPool(redis.Pool(), time.Minute)

	r.GET("/article/detail", cache.CachePage(store, time.Minute, func(c *gin.Context) {
		ArticleService.Detail(request.New(c))
	}))

}
