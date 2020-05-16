package env

import "time"

var (
	Server = struct {
		Port            int
		DebugMode       string
		ReadTimeout     int
		WriteTimeout    int
		ShutdownTimeout int
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
	MongoDBConfig = struct {
		Host          string
		Username      string
		Password      string
		LogCollection string
		Database      string
		Timeout       int
	}{}

	//----------- 定义配置文件结构体 -----------
	//ConfigName = struct {
	//	Key1 int
	//	Key2 string
	//	Keyn string
	//}{}
	//-----------------------------------------------

)
