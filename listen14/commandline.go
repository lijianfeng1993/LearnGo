package main

//命令行参数处理

import (
	"flag"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

//使用os包获取命令行参数
func test_os() {
	fmt.Println("args[0]:", os.Args[0])
	if len(os.Args) > 1 {
		for index, v := range os.Args {
			if index == 0 {
				continue
			}
			fmt.Printf("args[%d]: %v\n", index, v)
		}

	}
}

//使用flag包获取命令行参数
var recusive bool
var test string
var level int

func test_flag() {
	flag.BoolVar(&recusive, "r", false, "recusive ***")
	flag.StringVar(&test, "t", "default", "test ***")
	flag.IntVar(&level, "l", 1, "level ***")
	flag.Parse()
	fmt.Println(recusive, test, level)
}

//使用urfave/cli包的使用
// go get github.com/urfave/cli
// 获取命令行参数
func test_cli() {
	app := cli.NewApp()

	app.Name = "greet"
	app.Usage = "use cammandline."
	app.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args().Get(0)
		}
		fmt.Println("hello:", cmd)
		return nil
	}
	app.Run(os.Args)
}

func test_cli2() {
	var language string
	var recusive bool

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang, l",
			Value:       "english",
			Usage:       "select language",
			Destination: &language,
		},
		cli.BoolFlag{
			Name:        "recusive, r",
			Usage:       "recusive for the greeting",
			Destination: &recusive,
		},
	}

	app.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args().Get(0)
			fmt.Println("hello:", cmd)
		}
		fmt.Println("language: ", language)
		fmt.Println("recusive: ", recusive)
		return nil
	}

	app.Run(os.Args)
}

func main() {
	//test_os() //go run commandline.go a b c
	//test_flag()  //go run commmandline.go -r -t hello -l 5
	test_cli2()  // go run commandline.go hello -l china -r
}
