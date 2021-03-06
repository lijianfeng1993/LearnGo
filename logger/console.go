package logger

import (
	"os"

	"fmt"

)

type ConsoleLogger struct {
	level int
}

func NewConsoleLogger(config map[string]string) (log LogInterface, err error) {
	logLevel, ok := config["logLevel"]
	if !ok {
		err = fmt.Errorf("not fount logLevel")
	}
	level := getLogLevel(logLevel)

	log = &ConsoleLogger{
		level: level,
	}
	return
}

func (c *ConsoleLogger) Init(){
}

func (c *ConsoleLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		c.level = LogLevelDebug
	}
	c.level = level
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}){
	if c.level > LogLevelDebug {
		return
		}
	writeLog(os.Stdout, LogLevelDebug, format, args...)
}

func (c *ConsoleLogger) Trace(format string, args ...interface{}){
	if c.level > LogLevelTrace {
		return
	}
	writeLog(os.Stdout, LogLevelTrace, format, args...)
}

func (c *ConsoleLogger) Info(format string, args ...interface{}){
	if c.level > LogLevelInfo {
		return
	}
	writeLog(os.Stdout, LogLevelInfo, format, args...)
}

func (c *ConsoleLogger) Warn(format string, args ...interface{}){
	if c.level > LogLevelWarn {
		return
	}
	writeLog(os.Stdout, LogLevelWarn, format, args...)
}

func (c *ConsoleLogger) Error(format string, args ...interface{}){
	if c.level > LogLevelError {
		return
	}
	writeLog(os.Stdout, LogLevelError, format, args...)
}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}){
	if c.level > LogLevelFatal {
		return
	}
	writeLog(os.Stdout, LogLevelFatal, format, args...)
}

func (c *ConsoleLogger) Close(){

}