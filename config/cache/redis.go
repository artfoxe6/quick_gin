package cache

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"log"
	"quick_gin/config/env"
	"time"
)

var instance = new(redis.Pool)
var isLoad = false

// 建立连接
func connection() {
	c := env.Redis
	instance = &redis.Pool{
		Dial: func() (con redis.Conn, err error) {
			con, err = redis.Dial("tcp", c.Host,
				redis.DialPassword(c.Password),
				redis.DialDatabase(c.Db),
				redis.DialConnectTimeout(time.Second*time.Duration(c.Timeout)),
				redis.DialReadTimeout(time.Second*time.Duration(c.Timeout)),
				redis.DialWriteTimeout(time.Second*time.Duration(c.Timeout)))
			if err != nil {
				log.Fatalln(err.Error())
			}
			return con, err
		},
		MaxIdle:         c.MaxIdle,
		MaxActive:       c.MaxActive,
		IdleTimeout:     c.IdleTimeout,
		Wait:            true,
		MaxConnLifetime: 0,
	}
	isLoad = true
}

//获取实例
func Instance() redis.Conn {
	if !isLoad {
		connection()
	}
	return instance.Get()
}

func Pool() *redis.Pool {
	if !isLoad {
		connection()
	}
	return instance
}

//缓存字节数组到缓存
func Set(key string, data interface{}, ex int) {
	temp, _ := json.Marshal(data)
	_, err := Instance().Do("set", key, temp, "EX", ex)
	if err != nil {
		//log.Printf("%v", err)
	}
}

//从缓存中取数据
func Get(key string) (interface{}, bool) {
	var dst interface{}
	ext, _ := redis.Bool(Instance().Do("exists", key))
	if ext {
		d, _ := redis.Bytes(Instance().Do("get", key))
		err := json.Unmarshal(d, &dst)
		if err == nil {
			return dst, true
		} else {
			//log.Fatalf("%v", err)
			return nil, false
		}
	}
	return nil, false
}
