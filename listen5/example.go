package main

import (
	"fmt"
)

//求1-100之内的所有质数
func ifzhishu(n int) bool {
	if n <=1 {
		return false
	}
	for i:=2; i<n; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func zhishu() {
	for i:=1; i<100; i++ {
		if ifzhishu(i) == true {
			fmt.Println(i,"is zhishu\n")
		}
	}
}

//输入一行字符串，分别统计其中英文字母、空格、数字、和其他字符的个数
func calc(str string) (charCount int, numCount int, spaceCount int, otherCount int) {
	utfChars := []rune(str)
	for i:=0; i<len(utfChars); i++ {
		if utfChars[i] >= 'a' && utfChars[i] < 'z' || utfChars[i] >= 'A' && utfChars[i] <= 'Z'{
			charCount++
			continue
		}
		if utfChars[i] >= '0' && utfChars[i] <= '9' {
			numCount++
			continue
		}
		if utfChars[i] == ' ' {
			spaceCount++
			continue
		}
		otherCount++
	}
	return
}

func example(){
	var str string = "我爱北京a123  23ss 5"
	charCount, numCount, spaceCount, otherCount := calc(str)
	fmt.Println(charCount, numCount, spaceCount, otherCount)
}

func main(){
	//zhishu()
	example()
}
