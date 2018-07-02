package main

import (
	"fmt"
)

//每个变量都有内存地址，可以说通过变量的方式来操作对应大小的内存
func test1(){
	var a int32
	fmt.Printf("the addr of a:%p", &a)
}

//普通变量存储的是对应类型的值，这些类型就叫值类型
//指针类型的变量存储的是一个地址，所以又叫指针类型或引用类型
//指针类型定义，var 变量名 *类型
//每个变量都有自己的地址，但是指针变量里存储的也是一个地址
func test2(){
	a := 255
	fmt.Printf("The addr of a:%p, a:%d\n", &a, a)

	var b *int
	//b 是一个空指针，里面存的是空地址
	fmt.Printf("The addr of b:%p, b:%v\n", &b, b)

	b = &a  //把a的地址取出来，赋值给b
	fmt.Printf("The addr of b:%p, b:%v\n", &b, b)
}

//操作指针变量指向的地址里的值
//通过*符号可以获取指针变量指向的变量
func test3(){
	a := 200
	var b *int = &a
	fmt.Printf("b指向的地址存储的值为：%d\n", *b)

	//通过指针修改变量a的值
	*b = 1000
	fmt.Printf("b指向的地址存储的值为：%d\n", *b)
	fmt.Printf("a的值为：%d\n", a)
}

//指针变量传参
func modify(a *int){
	*a = 100
}
func test4(){
	var b int = 10
	p := &b
	modify(p)
	fmt.Printf("b:%d\n", b)
}

//指针变量传参2
func modify2(arr *[3]int){
	(*arr)[0] = 100
}
func test5(){
	a := [3]int{1,2,3}
	modify2(&a)
	//fmt.Println(a)
}

//注意：切片传参，切片是引用类型
func modify3(sls []int){
	sls[0] = 100
}
func test6(){
	a := [3]int{1,2,3}
	modify3(a[:])
	fmt.Println(a)
}

//make 和 new 的区别
//make 用来分配引用类型的内存，如map,slice,channel
//new 用来分配除引用类型的所有其他类型的内存，如int,数组等
func test7(){
	/* 会报错
	var a *int
	*a = 100
	*/
	var a *int = new(int)
	*a = 100
	fmt.Printf("*a=%d\n", *a)

	var b *[]int = new([]int)   //定义一个指针指向空切片
	fmt.Printf("*b=%v\n", *b)

	//(*b)[0] = 100
	//fmt.Printf("*b=%v\n", *b)  切片没有初始化，无法修改

	(*b) = make([]int,5,10)
	(*b)[0] = 100
	fmt.Printf("*b=%v\n", *b)
}

//值拷贝和引用拷贝
//值拷贝
func modifyInt(a int){
	a = 1000
}
func testmodifyInt(){
	var b int = 10
	modifyInt(b)
	fmt.Println(b)  //10
}
//引用拷贝
func modifyInt2(a *int){
	*a = 1000
}
func testmodifyInt2(){
	var a int = 100
	var b *int = &a
	modifyInt2(b)
	fmt.Println(a)
}

func main(){
	testmodifyInt2()
}
