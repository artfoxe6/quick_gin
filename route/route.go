package route

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"quick_gin/config/env"
	"quick_gin/route/api"
	"strings"
	"time"
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
	if env.ErrLog.Open == 1 {
		if strings.ToUpper(env.ErrLog.Type) == "LOCAL" {
			Route.Use(gin.RecoveryWithWriter(stdOutPut(env.ErrLog.LocalPath)))
		} else {
			Route.Use(gin.RecoveryWithWriter(sentryOutPut(env.ErrLog.SentryUrl)))
		}
	}

	//使用访问日志中间件
	if strings.ToUpper(env.Server.DebugMode) == "DEBUG" && env.StdLog.Open == 1 {
		Route.Use(gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: func(params gin.LogFormatterParams) string {
				return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
					params.ClientIP,
					params.TimeStamp.Format(time.RFC1123),
					params.Method,
					params.Path,
					params.Request.Proto,
					params.StatusCode,
					params.Latency,
					params.Request.UserAgent(),
					params.ErrorMessage,
				)
			},
			Output:    stdOutPut(env.StdLog.Path),
			SkipPaths: nil,
		}))
	}
	//其他路由
	api.LoadUserRoute(Route)
	api.LoadArticleRoute(Route)
	return Route
}

type stdOutPut string
type sentryOutPut string

func (o stdOutPut) Write(data []byte) (int, error) {
	logFile := string(o)
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

func (sentryOutPut) Write(data []byte) (int, error) {
	raven.CaptureMessage(string(data), map[string]string{"category": "RecoveryWithWriter"})
	return len(data), nil
}
