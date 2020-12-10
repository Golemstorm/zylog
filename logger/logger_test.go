package logger

import (
	"fmt"
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	InitTcpConnect("127.0.0.1", "12201", 2, time.Second*1)
	SetLogConfig("zylogdemo", "test", "", "v1.0", "./testlog")
	for i := 0; i < 100000; i++ {
		go Info(fmt.Sprintf("%v:test", i))
	}
	time.Sleep(time.Second * 120)
}
