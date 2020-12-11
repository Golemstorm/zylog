package logger

import (
	"encoding/json"
	"fmt"
	"github.com/Golemstorm/zylog/color"
	"github.com/Golemstorm/zylog/reverse"
	"io"
	"net"
	"os"
	"runtime/debug"
	"time"
)

var tcpClient *TcpClient
var logConfig *config
var connected bool

type TcpClient struct {
	Host       string
	Port       string
	tcpConn    *net.Conn
	Interval   time.Duration `json:"interval"`
	RetryTimes int           `json:"retry_times"`
}
type log struct {
	Level        int    `json:"level"`
	ShortMessage string `json:"short_message"`
	FullMessage  string `json:"full_message"`
	Timestamp    int64  `json:"timestamp"`
	Topic        string `json:"topic"`
	Type         string `json:"type"`
	Host         string `json:"host"`
	Version      string `json:"version"`
}
type config struct {
	Topic    string `json:"topic"`
	Type     string `json:"type"`
	Host     string `json:"host"`
	Version  string `json:"version"`
	FilePath string `json:"file_path"`
	switchs  bool   `json:"switchs"`
}

const (
	l_info = iota
	l_warn
	l_error
)

func (c *config) Switch(s bool) {
	c.switchs = s
}
func GetConfig() *config {
	return logConfig
}

const (
	log_service = "service_log"
	log_package = "package_log"
	log_debug   = "debug_log"
)

const (
	time_format      = "20060102"
	month_day_format = "0102"
	east_time        = "0229"
)

func Error(err error) {
	errmsg := err.Error()
	var log = log{
		Level:        l_error,
		Topic:        logConfig.Topic,
		Host:         logConfig.Host,
		Version:      logConfig.Version,
		Type:         logConfig.Type,
		ShortMessage: errmsg,
		FullMessage:  string(debug.Stack()) + errmsg,
		Timestamp:    time.Now().Unix(),
	}
	a, err := json.Marshal(log)
	if err != nil {
		fmt.Println(err)
		return
	}
	printLog(color.Red(logFormat("[error] "+string(debug.Stack())+errmsg, time.Now())), time.Now())
	sendLog(string(a))

}

func Warn(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	var log = log{
		Level:        l_warn,
		Topic:        logConfig.Topic,
		Host:         logConfig.Host,
		Version:      logConfig.Version,
		ShortMessage: msg,
		FullMessage:  string(debug.Stack()) + msg,
		Timestamp:    time.Now().Unix(),
	}
	a, err := json.Marshal(log)
	if err != nil {
		fmt.Println(err)
		return
	}

	printLog(color.Blue(logFormat("[warn] "+string(debug.Stack())+msg, time.Now())), time.Now())
	sendLog(string(a))

}

func Info(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	var log = log{
		Level:        l_info,
		Topic:        logConfig.Topic,
		Host:         logConfig.Host,
		Version:      logConfig.Version,
		ShortMessage: msg,
		FullMessage:  string(debug.Stack()) + msg,
		Timestamp:    time.Now().Unix(),
	}
	a, err := json.Marshal(log)
	if err != nil {
		fmt.Println(err)
		return
	}
	printLog(color.Green(logFormat("[info] "+string(debug.Stack())+msg, time.Now())), time.Now())
	sendLog(string(a))
}

func getTcpClient() *TcpClient {
	return tcpClient
}

func setTcpClient(host, port string, interval time.Duration, retrytime int) {
	tcpClient.Host = host
	tcpClient.Port = port
	tcpClient.Interval = interval
	tcpClient.RetryTimes = retrytime
}

func tcpClientConnect() {
	con, err := net.Dial("tcp", fmt.Sprintf("%v:%v", tcpClient.Host, tcpClient.Port))
	if err != nil {
		fmt.Println(err)
		connected = false
		return
	}
	tcpClient.tcpConn = &con
	connected = true
}

func checkConnected() {
	if !connected {
		tcpClientConnect()
	}
}

func sendLog(msg string) {
	writeLog([]byte(msg), 0)
}

func writeLog(bys []byte, depth int) {
	if depth > 2 {
		fmt.Println("超出重试次数")
		connected = false
		err := writeFile(fmt.Sprintf("%v.log", time.Now().Format(time_format)), string(bys))
		fmt.Println(err)
		//todo
		return
	}
	if connected {
		cons := *tcpClient.tcpConn
		if cons == nil {
			connected = false
			time.Sleep(tcpClient.Interval + time.Nanosecond*100)
			writeLog(bys, depth+1)
			return
		}
		a := append(bys, 0)
		_, err := cons.Write(a)
		if err != nil {
			connected = false
			time.Sleep(tcpClient.Interval + time.Nanosecond*100)
			writeLog(bys, depth+1)
			return
		}
	} else {
		connected = false
		time.Sleep(tcpClient.Interval + time.Nanosecond*100)
		writeLog(bys, depth+1)
		return
	}
}

func SetLogConfig(topic, types, host, version, logpath string) {
	if host == "" {
		logConfig.Host = getLocalIP()
	} else {
		logConfig.Host = host
	}
	if logpath != "" {
		logConfig.FilePath = logpath
	} else {
		logConfig.FilePath = "./misslog"
	}
	logConfig.Topic = topic
	logConfig.Type = types
	logConfig.Version = version
	logConfig.switchs = false

}

func InitTcpConnect(host, port string, RetryTimes int, intervals ...time.Duration) {
	tcpClient = new(TcpClient)
	logConfig = new(config)
	connected = false
	interval := time.Second * 1
	if len(intervals) > 0 {
		interval = intervals[0]
	}
	retrytime := 10
	if RetryTimes > 0 {
		retrytime = RetryTimes
	}
	setTcpClient(host, port, interval, retrytime)
	checkConnected()
	go func() {
		for {
			checkConnected()
			time.Sleep(interval)
		}
	}()
}

func getLocalIP() (ipv4 string) {
	var (
		addrs   []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
	)
	// 获取所有网卡
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	// 取第一个非lo的网卡IP
	for _, addr = range addrs {
		// 这个网络地址是IP地址: ipv4, ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// 跳过IPV6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String() // 192.168.1.1
				return
			}
		}
	}
	return
}

func createDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	os.Chmod(path, os.ModePerm)
	return nil
}

func isExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}

func writeFile(fileName, msg string) error {
	if !isExist(logConfig.FilePath) {
		if err := createDir(logConfig.FilePath); err != nil {
			return err
		}
	}
	var (
		err error
		f   *os.File
	)

	f, err = os.OpenFile(logConfig.FilePath+"/"+fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	_, err = io.WriteString(f, msg+"\r\n")

	defer f.Close()
	return err
}

func logFormat(msg string, times time.Time) string {
	return fmt.Sprintf("%v %v", times.Format(time.RFC3339), msg)
}

func printLog(msg string, times time.Time) {
	if times.Format(month_day_format) == east_time && logConfig.switchs {
		fmt.Println(reverse.Reverse(msg))
	} else {
		fmt.Println(msg)
	}
}
