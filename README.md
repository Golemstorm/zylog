# zylog
worst golang log package ever

## how to use
1.first you need import this logger package,then init the tcp connection. 
```
 logger.InitTcpConnect("127.0.0.1", "12201", 2, time.Second*1)   
 
                      //host//port//retry times//interval time 
```
                       
2.set the log config
```
SetLogConfig("topic", "tyoe", "source_host", "version", "./fail_log_path")
```

3.use logger

```
logger.Info("string")
logger.Error(errors.new("error"))
logger.Warn("string")
```