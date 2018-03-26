package until

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
	logger "until/xlog4go"
)

func Assert(condition bool, msg string, v ...interface{}) {
	if !condition {
		panic(fmt.Sprintf("assertion failed: "+msg, v...))
	}
}
func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
func Memery() {
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		logger.Info("Alloc = %vMB  TotalAlloc = %vMB  Sys = %vMB  NumGC = %vMB\n", m.Alloc/(1024*1024), m.TotalAlloc/(1024*1024), m.Sys/(1024*1024), m.NumGC)
		time.Sleep(5 * time.Second)
	}
}
