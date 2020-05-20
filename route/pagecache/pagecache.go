package pagecache

import (
	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
	envRedis "quick_gin/config/cache"
	"time"
)

//接口缓存多少秒
func Second(second int, handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return cache.CachePage(
		persistence.NewRedisCacheWithPool(envRedis.Pool(), time.Minute),
		time.Duration(second)*time.Second,
		handlerFunc)
}

//接口缓存5分钟
func FiveMinute(handlerFunc gin.HandlerFunc) gin.HandlerFunc {
	return cache.CachePage(
		persistence.NewRedisCacheWithPool(envRedis.Pool(), time.Minute),
		5*time.Minute,
		handlerFunc)
}
