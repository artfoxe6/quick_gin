package lib

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"os"
	"quick_gin/config/env"
	"strconv"
	"strings"
)

func Int(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		ErrLog(e)
		return 0
	}
	return i
}

func String(num int) string {
	return strconv.Itoa(num)
}

func ErrLog(err error) {
	if err == nil {
		return
	}
	if env.ErrLog.Open != 1 {
		return
	}
	if strings.ToUpper(env.ErrLog.Type) == "SENTRY" {
		raven.CaptureError(err, nil)
	} else {
		logFile := string(env.ErrLog.LocalPath)
		f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			return
		}
		defer f.Close()
		f.Write([]byte(fmt.Sprintf("%v", err)))
	}
}

// 将m2合并到m1
func MergeMap(m1, m2 map[string]interface{}) {
	for k, v := range m2 {
		m1[k] = v
	}
}
