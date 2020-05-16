package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"quick_gin/config/env"
	"quick_gin/route/api"
	"quick_gin/util/qlog"
)

var Route *gin.Engine

func Load() *gin.Engine {
	//设置运行模式
	gin.SetMode(env.Server.DebugMode)

	//新建一个空路由
	Route = gin.New()

	//设置同源策略 默认允许所有
	Route.Use(cors.Default())

	//程序错误时恢复运行 并记录日志
	Route.Use(gin.RecoveryWithWriter(qlog.QLog{}))

	//开启访问日志中间件
	//if strings.ToUpper(env.Server.DebugMode) == "DEBUG" {
	//	Route.Use(gin.LoggerWithConfig(gin.LoggerConfig{
	//		Formatter: func(params gin.LogFormatterParams) string {
	//			return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//				params.ClientIP,
	//				params.TimeStamp.Format(time.RFC1123),
	//				params.Method,
	//				params.Path,
	//				params.Request.Proto,
	//				params.StatusCode,
	//				params.Latency,
	//				params.Request.UserAgent(),
	//				params.ErrorMessage,
	//			)
	//		},
	//		Output:    qlog.QLog{},
	//		SkipPaths: nil,
	//	}))
	//}
	Route.Static("/apidoc", "apidoc/")
	//其他路由
	api.LoadUserRoute(Route)
	api.LoadArticleRoute(Route)
	return Route
}
