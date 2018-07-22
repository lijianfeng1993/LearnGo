package main

import "LearnGo/listen17/log"

func main(){
	file := log.NewFileLog("e:/log.txt")
	file.LogDebug("this is Debug log.")
	file.LogWarn("this is Warn log.")
}
