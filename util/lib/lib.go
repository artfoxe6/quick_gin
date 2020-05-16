package lib

import (
	"strconv"
)

func Int(s string) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		return 0
	}
	return i
}
