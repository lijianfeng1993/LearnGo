package main

import (
	"fmt"
)

func main(){
	/*
	const a int = 100
	const b string = "hello"
	fmt.Printf("a=%d\n", a)
	fmt.Printf("b = %s\n", b)
	*/
	//优雅的写法
	const (
		a int = 100
		b string = "hello"
		c   //不写的话默认为和上一个一样
	)
	fmt.Printf("a=%d\n", a)
	fmt.Printf("b = %s\n", b)
	fmt.Printf("c = %s\n", c)

	const (
		e = iota //iota默认值为0，没增加一行，加1
		f
		g = iota
	)
	fmt.Println(e,f,g)

	const (
		a1 = 1 << iota
		a2
		a3
	)
	fmt.Println(a1, a2, a3)

	const (
		A = iota
		B
		C
		D = 8
		E
		F = iota
		G
	)
	fmt.Println(A,B,C,D,E,F,G)
}
