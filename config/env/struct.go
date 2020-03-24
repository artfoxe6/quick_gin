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

	//----------- 定义配置文件结构体 -----------
	//ConfigName = struct {
	//	Key1 int
	//	Key2 string
	//	Keyn string
	//}{}
	//-----------------------------------------------

)
