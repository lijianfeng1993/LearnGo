package main

import (
	"github.com/urfave/cli"
	"os"
	"bufio"

	"strings"
	"fmt"


)

//读取计算表达式
func getInput() (string,error){
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func process(c *cli.Context) (err error) {
	for {
		express, _ := getInput()
		if len(express) == 0 {
			continue
		}
		express = strings.TrimSpace(express)
		postexpress, errRet := transPostExpress(express)  //转为后缀表达式
		if errRet != nil {
			err = errRet
			fmt.Println(err)
			return
		}

		//fmt.Println(postexpress)
		result, err := calc(postexpress)
		if errRet != nil {
			err = errRet
			fmt.Println("error:",err)
			continue
		}
		fmt.Println(result)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "calculate"
	app.Usage = "calculate express"

	app.Action = func(c *cli.Context) error {
		return process(c)
	}
	app.Run(os.Args)
}
