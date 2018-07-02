package main

import (
	"fmt"
)

//defer 定义的不会立即执行，在函数返回之前执行
func testDefer1(){
	defer fmt.Println("hello")
	fmt.Println("aaa")
	fmt.Println("bbb")
}


//defer 按照栈的顺序，先输出helloworld，然后在输出hello
func testDefer2(){
	defer fmt.Println("hello")
	defer fmt.Println("helloworld")
	fmt.Println("aaa")
	fmt.Println("bbb")
}

//4,3,2,1,0
func testDefer3(){
	for i := 0; i<5; i++ {
		defer fmt.Println(i)
	}
}


func main(){
	testDefer1()
}