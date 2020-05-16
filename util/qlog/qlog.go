package qlog

import (
	"encoding/json"
	"io"
	"os"
)

type QLog struct {
}

func (log QLog) Write(data []byte) (int, error) {
	logString := string(data)
	return WriteLog(logString)
}

func WriteLog(v interface{}) (int, error) {
	log, err := json.Marshal(v)

	if err != nil {
		return 0, err
	}
	//默认将日志打印到标准输出,将日志存到mongodb或者es等请自行实现
	return io.WriteString(os.Stdout, string(log))
}
