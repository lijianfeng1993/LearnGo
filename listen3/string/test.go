package main

import (
	"fmt"
)

//字符串逆序
func testReverseString(){
	var str = "hello"
	var bytes []byte = []byte(str)

	for i :=0; i<len(str)/2; i++ {
		tmp := bytes[len(str)-i-1]
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = tmp
	}

	str = string(bytes)
	fmt.Println(str)
}

//对包含中文的字符串进行逆序
func testReverseStringv2(){
	var str = "hello你好"
	var r []rune = []rune(str)

	for i := 0; i<len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}

	str = string(r)
	fmt.Println(str)
}

//判断字符串是否是回文
func testHuiwen(){
	var str = "上海自来水来看字海上"
	var r []rune = []rune(str)

	for i := 0; i<len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}

	str2 := string(r)
	if str == str2{
		fmt.Println("is huiwen")
	}else{
		fmt.Println("not huiwen")
	}
}

func main(){
	testReverseString()
	testReverseStringv2()
	testHuiwen()
}
