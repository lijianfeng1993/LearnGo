package main

import (
	"fmt"
	"time"
)

//写一个程序，统计一段代码的执行时间，单位精确到微秒
func testCostTime(){
	start := time.Now().UnixNano()
	for i :=0 ; i<10; i++{
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	cost := (end - start)/1000
	fmt.Printf("Code cost:%d us\n",cost)
}



func main(){
	testCostTime()
}
