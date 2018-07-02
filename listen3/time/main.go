package main

import (
	"time"
	"fmt"

)

func testTime(){
	now := time.Now()
	fmt.Println(now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minite := now.Minute()
	send := now.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minite,send)

	//时间戳
	timestamp := now.Unix()
	fmt.Printf("timestamp is : %d\n", timestamp)


}

//时间戳转换成当前时间
func testTimestamp(timestamp int64){
	timeObj := time.Unix(timestamp,0)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minite := timeObj.Minute()
	send := timeObj.Second()

	fmt.Printf("current timestamp:%d\n", timestamp)
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n",year,month,day,hour,minite,send)

}

//定时器使用,每两秒打印一下时间
func testTicker(){
	ticker := time.Tick(2*time.Second)
	for i := range ticker {
		fmt.Printf("%v\n",i)
	}
}

//格式化,方法2
func testFormat(){
	now := time.Now()
	//以go诞生的时间为模板进行格式化
	fmt.Println(now.Format("2006/01/02 15:04:05"))
}


func main(){
	testTime()
	timestamp := time.Now().Unix()
	testTimestamp(timestamp)
	//testTicker()
	testFormat()
}
