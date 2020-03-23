package env

import (
	"github.com/getsentry/raven-go"
	"github.com/go-ini/ini"
	"log"
	"strings"
)

var (
	Server = struct {
		Port            int
		DebugMode       string
		ReadTimeout     int
		WriteTimeout    int
		ShutdownTimeout int
	}{}
	ErrLog = struct {
		Open      int
		Type      string
		LocalPath string
		SentryUrl string
	}{}
	StdLog = struct {
		Open int
		Path string
	}{}

	//----------- (1)第一步，定义配置文件映射 -----------
	//ConfigName = struct {
	//	Key1 int
	//	Key2 string
	//	Keyn string
	//}{}
	//-----------------------------------------------

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

	//---------- (2)第二步，配置文件到结构体的映射 -----------
	//
	//mapTo(h, "configName", &ConfigName)
	//---------------------------------------

	if strings.ToUpper(ErrLog.Type) == "SENTRY" && ErrLog.Open == 1 {
		initSentry(ErrLog.SentryUrl)
	}
}

func mapTo(h *ini.File, section string, v interface{}) {
	if err := h.Section(section).MapTo(v); err != nil {
		log.Fatal(err.Error())
	}
}

func initSentry(url string) {
	err := raven.SetDSN(url)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
