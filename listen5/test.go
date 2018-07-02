package main

import "fmt"

//可变参数
func calc(b ...int) int {
	sum := 0
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}
	return sum
}

func calc_2(a int,b ...int) int {
	sum := 0
	for i := 0; i < len(b); i++ {
		sum = sum + b[i]
	}
	return sum
}
func main(){
	sum := calc()
	sum1 := calc(10)
	sum2 := calc(10,20,30)
	fmt.Println(sum, sum1, sum2)

	sum3 := calc_2(1)
	sum4 :=calc_2(1,20)
	fmt.Println(sum3,sum4)
}