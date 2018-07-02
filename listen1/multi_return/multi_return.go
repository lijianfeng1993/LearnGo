package main

import (
	"fmt"
)


func add(a int, b int) (int,int) {
	return a + b, 5
}

func main() {
	sum, a := add(1,2)
	fmt.Println(sum, a)
}