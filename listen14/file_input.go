package main

import (
	"fmt"
	"os"
)

func main(){
	var a int
	var b string
	var c float32
	fmt.Fscanf(os.Stdin, "%d%s%f", &a, &b, &c)
	//fmt.Scanf("%d%s%f", &a, &b, &c)   Scanf 是对Fscanf的封装
	//fmt.Println(a,b,c)
	fmt.Fprintln(os.Stdout, a, b, c)
}
