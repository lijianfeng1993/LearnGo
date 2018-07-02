package main

import (
	"fmt"
	"time"
)

//break 终止当前整个循环
//continue 终止本次循环


//无限循环
func test(){
	for {
		fmt.Println("hello")
		time.Sleep(time.Second)
	}
}

