package main

import (
	"LearnGo/logger"
	"time"
)



func initLogger(name, logPath, logName string, level string) (err error){
	m := make(map[string]string, 8)
	m["logPath"] = logPath
	m["logName"] = logName
	m["logLevel"] = level
	err = logger.InitLogger(name, m)
	if err != nil {
		return
	}
	//log = logger.NewConsoleLogger(level)
	logger.Debug("init logger success.")
	return
}

func Run(){
	for {
		logger.Debug("user server is running")
		time.Sleep(time.Second)
	}
}

func main(){
	initLogger("console","E:/GoProject/src/LearnGo/Listen17/log", "filetest", "debug")
	Run()
	return
}