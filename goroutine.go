package reentrantlock

import (
	"runtime"
	"strconv"
	"strings"
)

func goroutineId() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	id, _ := strconv.Atoi(strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0])
	return id
}
