package main

import (
	"fmt"
	"strings"

)

func testStr(){
	var a string = "hello"
	var b  = "world"
	c := "lijianfeng"
	fmt.Println(a,b,c)

	clen := len(c)
	fmt.Printf("len of c = %d\n", clen)

	//字符串拼接
	c = c + c
	fmt.Println(c)

	b = fmt.Sprintf("%s%s",b,b)
	fmt.Println(b)

	//字符串分割
	ips := "10.0.1.1;10.0.2.1"
	ipSlice := strings.Split(ips,";")
	fmt.Println(ipSlice[0])
	fmt.Println(ipSlice[1])

	//是否包含
	result := strings.Contains(ips,"10.0.1.1")
	fmt.Println(result)

	//前缀,后缀,位置
	str := "http://www.baidu.com"
	if strings.HasPrefix(str, "http"){
		fmt.Println("yes")
	}else {
		fmt.Println("no")
	}

	if strings.HasSuffix(str,"com"){
		fmt.Println("yes")
	}else {
		fmt.Println("no")
	}

	index := strings.Index(str,"baidu") //first index
	fmt.Println(index)

	last_index := strings.LastIndex(str,"o")
	fmt.Println(last_index)

	//join操作
	var strlist []string = []string{"10.0.1.1","10.0.2.1","10.0.3.1"}
	resultStr := strings.Join(strlist,";")
	fmt.Println(resultStr)
}

func main(){
	testStr()
}
