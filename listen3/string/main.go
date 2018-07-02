package main

import (
	"fmt"
)

func testString(){
	//字符串底层是一个byte数组，可以和[]byte类型相互转换
	var str = "hello"
	fmt.Printf("str[0] = %c\n", str[0])

	for index, val := range str {
		fmt.Printf("str[%d] = %c\n", index, val)
	}

	//字符串只读，无法直接修改，需要先将其转化成切片,单引号是字符，双引号是字符串
	var byteSlice []byte
	byteSlice = []byte(str)
	byteSlice[0] = 'w'
	str = string(byteSlice)
	fmt.Println(str)

	//字符串是由Byte字节组成的，所以字符串的长度就是byte字节的长度,一个中文占3个字节
	fmt.Printf("len(str) = %d\n", len(str))

	str = "hello,李剑锋"
	fmt.Printf("len(str) = %d\n", len(str))

	//rune类型用来表示utf8字符，一个rune字符由1个或多个byte组成
	var b rune = '中'
	fmt.Printf("b = %c\n", b)

	var runeSlice []rune
	runeSlice = []rune(str)
	fmt.Printf("str 长度 = %d, len(str) = %d\n", len(runeSlice), len(str))
}

func main(){
	testString()
}