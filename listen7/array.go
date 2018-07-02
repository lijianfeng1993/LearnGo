package main

import(
	"fmt"
	"rand"
	"time"
)

//数组
func testArray(){
	var a [5]int
	a[0] = 200
	a[1] = 300
	fmt.Println(a)

	var b [5]int = [5]int{1,2,3,4,5}
	fmt.Println(b)

	c := [...]int{1,2,3,4,5,6}
	fmt.Println(c)

	d := [5]int{1,2,3}
	fmt.Println(d)   //1,2,3,0,0

	e := [5]int{3:100,4:200}
	fmt.Println(e)   //0,0,0,100,200

}

//数组遍历
func testArray2(){
	a := [5]int{3:100,4:300}
	for i:=0; i<len(a); i++ {
		fmt.Println(a[i])
	}

	for index,val := range a {
		fmt.Println(index, val)
	}
}

//二维数组
func testArray3(){

	var a [3][2]int
	a[0][0] = 10
	a[0][1] = 20
	a[1][0] = 30
	a[1][1] = 40
	a[2][0] = 50
	a[2][1] = 60
	fmt.Println(a)

	for i:=0; i<3; i++ {
		for j :=0; j<2; j++ {
			fmt.Println(a[i][j])
		}
		fmt.Println()
	}

	for i,val := range a{
		for j, value := range val {
			fmt.Printf("(%d,%d)=%d",i,j,value)
		}
	}

	b := [3][2]string{
		{"a","b"},
		{"c","d"},
		{"e","f"},
	}
	fmt.Println(b)
}

//数组拷贝
func testArray4(){
	a := [3]int{10,20,30}
	b := a
	b[0] = 1000
	fmt.Printf("a=%v\n",a)
	fmt.Printf("b=%v\n",b)
}

//数组是值类型，传参
func modify(b [3]int){
	b[0] = 1000
	fmt.Println(b)  //{1000,20,30}
}
func testArray5(){
	a := [3]int{10,20,30}
	modify(a)   //传进去的时候，会拷贝一份，不会影响a
	fmt.Println(a) //{10,20,30}
}

//求数组中所有元素之和
func testArray6(){
	a := [3]int{10,20,30}
	sum := 0
	for _,val := range a {
		sum += val
	}
	fmt.Println(sum)

	/*
	//初始化随机数种子
	rand.Seed(time.Now().Unix())
	var b [10]int
	for i:=0; i<len(b); i++ {
		//产生一个0--1000的随机数
		b[i] = rand.Intn(1000)
	}
	*/
}

//找出数组中和为给定值的两个元素的下标，如数组【1，3，5，8，7】，找出两个元素之和等于8的下标分别是（0，4）和（1，2）
func testArray7(){
	a := [5]int{1,3,5,8,7}
	for i:=0; i<len(a)-1; i++ {
		for j:=i+1; j<len(a); j++ {
			if a[i] + a[j] == 8 {
				fmt.Println(i,j)
			}
		}
	}
}

func main(){
	testArray7()
}