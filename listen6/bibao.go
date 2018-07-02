package main

import(
	"fmt"
	"strings"
)

//闭包：一个函数和与其相关的引用环境组合而成的实体
func Adder() func(int)int {
	var x int  //x初始化0
	return func(d int)int{
		x += d
		return x
	}
}
func testClosure1(){
	f := Adder()  //返回了一个匿名函数，x从0开始，只要f一直在调用，x就一直存活
	ret := f(1)   //最后x=1并返回
	fmt.Printf("ret = %d\n", ret)

	ret = f(20)   //x=21
	fmt.Printf("ret = %d\n", ret)

	ret = f(100)  //x=121
	fmt.Printf("ret = %d\n", ret)

	f1 := Adder()  //返回匿名函数，x从0开始，重新初始化
	ret = f1(1)
	fmt.Printf("ret = %d\n", ret)
}

//闭包例子2
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name+suffix
		}
		return name
	}
}
func testSuffix(){
	func1 := makeSuffixFunc(".jpg")
	func2 := makeSuffixFunc(".doc")
	fmt.Println(func1("a.jpg"))
	fmt.Println(func2("bbb"))
}


//闭包例子3
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base = base + i
		return base
	}
	sub := func(i int) int {
		base = base -i
		return base
	}
	return add, sub
}
func testCalc(){
	add, sub := calc(5)
	fmt.Println(add(1))
	fmt.Println(sub(2))
}

func main(){
	//testClosure1()
	//testSuffix()
	testCalc()
}