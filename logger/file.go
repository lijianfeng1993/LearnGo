package logger

import (
	"os"
	"fmt"
	"time"
	"path"
)

// 2018/3/26 0:01.383  DEBUG  logDebug.go:29  this is a debug log

type FileLogger struct {
	level int
	logPath string
	logName string
	file *os.File   //正常日志的文件句柄
	warnFile *os.File   //错误日志的文件句柄
}

func NewFileLogger(level int, logPath string, logName string) (LogInterface) {
	logger := &FileLogger{
		level:level,
		logPath:logPath,
		logName:logName,
	}
	logger.init()
	return logger
}

func (f *FileLogger) init() {
	filename := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err:%v", filename, err))
	}
	f.file = file

	//写错误日志和fatal日志的文件
	filename = fmt.Sprintf("%s/%s-error.log", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed, err:%v", filename, err))
	}
	f.warnFile = file
}

func (f *FileLogger) SetLevel(level int){
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) writeLog(file *os.File, level int, format string, args ...interface{}){
	if (f.level) > level {
		return
	}

	now := time.Now()
	nowStr := now.Format("2006-01-02 15:04:05.999")  //以go的诞生之日为时间点格式化，固定的，不能变

	levelStr := getLevelText(level)

	fileName, funcName, lineNo := GetLineInfo()
	fileName = path.Base(fileName)
	funcName = path.Base(funcName)

	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(file, "%s %s (%s:%s:%d) %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
}

func (f *FileLogger) Debug(format string, args ...interface{}){
	if f.level > LogLevelDebug {
		return
	}
	writeLog(f.file, LogLevelDebug, format, args...)
}

func (f *FileLogger) Trace(format string, args ...interface{}){
	if f.level > LogLevelTrace {
		return
	}
	writeLog(f.file, LogLevelTrace, format, args...)
}

func (f *FileLogger) Info(format string, args ...interface{}){
	if f.level > LogLevelInfo {
		return
	}
	writeLog(f.file, LogLevelInfo, format, args...)
}

func (f *FileLogger) Warn(format string, args ...interface{}){
	if f.level > LogLevelWarn {
		return
	}
	writeLog(f.warnFile, LogLevelWarn, format, args...)
}

func (f *FileLogger) Error(format string, args ...interface{}){
	if f.level > LogLevelError {
		return
	}
	writeLog(f.warnFile, LogLevelError, format, args...)
}

func (f *FileLogger) Fatal(format string, args ...interface{}){
	if f.level > LogLevelFatal {
		return
	}
	writeLog(f.warnFile, LogLevelFatal, format, args...)
}

func (f *FileLogger)Close(){
	f.file.Close()
	f.warnFile.Close()
}