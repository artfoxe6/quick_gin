package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"quick_gin/config/env"
)

var instance = new(gorm.DB)
var isLoad = false

// 建立连接
func connection() {
	var err error
	c := env.Database
	instance, err = gorm.Open(c.Connection, c.User+":"+c.Password+"@tcp("+c.Host+")/"+
		c.DbName+"?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("MySQL连接错误", err.Error())
	}
	isLoad = true
}

//获取连接实例
func Instance() *gorm.DB {
	if !isLoad {
		connection()
	}
	return instance
}
