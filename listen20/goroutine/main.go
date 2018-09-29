package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
goroutine 使用
*/
func hello(i int) {
	fmt.Println("hello,goroutine: ", i)
}
func test1() {
	//go hello()  //单独创建一个线程，主线程结束了，子线程可能还没创建好

	//多个线程并发执行
	for i := 0; i < 10; i++ {
		go hello(i)
	}

	fmt.Println("main thread terminate.")
	time.Sleep(time.Second)
}

/*
多核控制：
	通过runtime包进行多核设置
	GOMAXPROCS设置当前程序运行时占用的CPU核数
	NumCPU获取当前系统的CPU核数
*/
var i int

func calc() {
	for {
		i++
	}
}
func test2() {
	// 任务管理器，性能，查看cpu占用
	cpu := runtime.NumCPU()
	fmt.Println("cpu: ", cpu)
	//如果不设置cpu，就会把cpu跑满
	runtime.GOMAXPROCS(2)

	for i := 0; i < 10; i++ {
		go calc()
	}
	time.Sleep(time.Hour)
}

/*
Goroutine 原理：
	一个操作系统线程对应用户态多个goroutine
	同时使用多个操作系统线程
	操作系统线程对goroutine是多对多的关系

*/

/*
channel介绍
 	本质上就是一个队列，是一个容器
 	因此定义的时候，需要指定容器中元素的类型
 	var 变量名 chan 数据类型

元素入队和出列
	var a chan int
	入队： a <- 100
	出队： data := <- a
*/
func produce(c chan int){
	c <- 1000
	fmt.Println("produce finished.")
}
func consume(c chan int) {
	data := <- c
	fmt.Println(data)
}
func test3() {
	var c chan int
	fmt.Printf("c=%v\n", c)

	c = make(chan int, 10)  //初始化队列，容量为10
	fmt.Printf("c=%v\n", c)

	c <- 100
	//data := <- c
	//fmt.Println("data: ", data)

	//出队不进行操作，丢弃该元素
	<-c

	//不带缓冲器的队列
	var a chan int
	a = make(chan int)   //没有空间，插不进去
	go produce(a)
	go consume(a)
	time.Sleep(time.Second*5)
}


func main() {
	//test1()
	//test2()
	test3()
}
