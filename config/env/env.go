package env

import (
	"github.com/go-ini/ini"
	"log"
)

// 加载配置信息
func Load() {
	h, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("%v", err)
	}
	mapTo(h, "server", &Server)
	mapTo(h, "database", &Database)
	mapTo(h, "redis", &Redis)
	mapTo(h, "jwt", &Jwt)
	mapTo(h, "mongodb", &MongoDBConfig)

	//---------- 配置文件到结构体的映射 -----------
	//
	//mapTo(h, "configName", &ConfigName)
	//---------------------------------------
}

func mapTo(h *ini.File, section string, v interface{}) {
	if err := h.Section(section).MapTo(v); err != nil {
		log.Fatal(err.Error())
	}
}
