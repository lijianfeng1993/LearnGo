package main

import(
	"fmt"
)

//函数也是一种类型，因此可以定义一个函数类型的变量
func add(a,b int) int {
	return a+b
}
func testFunc1(){
	f1 := add
	fmt.Printf("type of f1:%T\n", f1)
	sum := f1(2,5)
	fmt.Println(sum)
}
//匿名函数使用
func testFunc2(){
	f1 := func (a,b int) int {
		return a+b
	}
	fmt.Printf("type of f1:%T\n", f1)
	sum := f1(2,3)
	fmt.Println(sum)
}

//定义defer的时候，i=0首先压栈，所以最后先打印100，再打印0
func testFunc3(){
	var i int = 0
	defer fmt.Printf("i=%d\n", i)
	i = 100
	fmt.Printf("i=%d\n", i)
	return
}
//defer中使用匿名函数,defer中调用了匿名函数,当程序最后执行defer的时候，i是100,所以最后打印两次100
func testFunc4(){
	var i int = 0
	defer func() {
		fmt.Printf("i=%d\n", i)
	}()
	i = 100
	fmt.Printf("i=%d\n", i)
	return
}

//函数作为一个参数
func sub(a,b int) int {
	return a-b
}

func calc(a,b int, op func(int,int) int) int {
	return op(a,b)
}
func testFunc5(){
	sum := calc(100,40, add)
	sub := calc(100,40, sub)
	fmt.Println(sum, sub)
}


func main(){
	//testFunc1()
	//testFunc2()
	testFunc5()
}