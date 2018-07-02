package main

import (
	"fmt"
	"bufio"
	"os"
)

// 从终端读取带空格的字符串

var inputReader *bufio.Reader
var input string
var err error

func main(){
	var str string
	reader := bufio.NewReader(os.Stdin)
	str,_ =reader.ReadString('\n')  //读到换行为止
	fmt.Println(str)
}