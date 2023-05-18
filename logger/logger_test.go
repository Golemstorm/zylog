package logger

import (
	"testing"
	"time"
)

func TestLogger(t *testing.T) {
	//InitTcpConnect("127.0.0.1", "12201", 2, time.Second*1)
	SetLogConfig("zylogdemo", "test", "", "v1.0", "./testlogs")
	for i := 0; i < 10; i++ {
		go Info("test %v", i)
	}
	time.Sleep(time.Second * 120)
}
