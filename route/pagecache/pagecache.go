package pagecache

import (
	envRedis "quick_gin/config/cache"
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	"time"
)

func Second(n int, handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return cache.CachePage(
		persistence.NewRedisCacheWithPool(envRedis.Pool(), time.Minute),
		time.Duration(n)*time.Second,
		handlerFunc)
}

func FiveMinute(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return cache.CachePage(
		persistence.NewRedisCacheWithPool(envRedis.Pool(), time.Minute),
		5*time.Minute,
		handlerFunc)
}
