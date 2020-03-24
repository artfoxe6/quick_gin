package env

import (
	"github.com/getsentry/raven-go"
	"github.com/go-ini/ini"
	"log"
	"strings"
	"time"
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
	Database = struct {
		Connection string
		User       string
		Password   string
		Host       string
		DbName     string
	}{}
	Redis = struct {
		Host        string
		Password    string
		MaxIdle     int
		MaxActive   int
		IdleTimeout time.Duration
		Db          int
		Timeout     int
	}{}
	Jwt = struct {
		Key string
		Exp int
	}{}

	//----------- (1)第一步，定义配置文件结构体 -----------
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
	mapTo(h, "database", &Database)
	mapTo(h, "redis", &Redis)
	mapTo(h, "jwt", &Jwt)

	//---------- (2)第二步，配置文件到结构体的映射 -----------
	//
	//mapTo(h, "configName", &ConfigName)
	//---------------------------------------

	if strings.ToUpper(ErrLog.Type) == "SENTRY" && ErrLog.Open == 1 {
		err := raven.SetDSN(ErrLog.SentryUrl)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}

func mapTo(h *ini.File, section string, v interface{}) {
	if err := h.Section(section).MapTo(v); err != nil {
		log.Fatal(err.Error())
	}
}
