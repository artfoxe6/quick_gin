package env

import (
	"github.com/getsentry/raven-go"
	"github.com/go-ini/ini"
	"log"
	"strings"
)

// 加载配置信息
func Load() {
	h, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("%v", err)
	}
	mapTo(h, "server", &Server)
	mapTo(h, "errLog", &ErrLog)
	mapTo(h, "stdLog", &StdLog)
	mapTo(h, "database", &Database)
	mapTo(h, "redis", &Redis)
	mapTo(h, "jwt", &Jwt)

	//---------- 配置文件到结构体的映射 -----------
	//
	//mapTo(h, "configName", &ConfigName)
	//---------------------------------------

	if strings.ToUpper(ErrLog.Type) == "SENTRY" && ErrLog.Open == 1 {
		err := raven.SetDSN(ErrLog.SentryUrl)
		if err != nil {
			log.Fatalln("Sentry错误", err.Error())
		}
	}
}

func mapTo(h *ini.File, section string, v interface{}) {
	if err := h.Section(section).MapTo(v); err != nil {
		log.Fatal(err.Error())
	}
}
