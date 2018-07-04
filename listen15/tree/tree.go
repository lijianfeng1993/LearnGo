package main

import (
	"github.com/urfave/cli"
	"io/ioutil"
	"fmt"
	"path/filepath"
	"os"
)

//实现一个类似linux tree命令，能够以树状形式显示目录文件


func ListDir(dirPath string, deep int) (err error){
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	if deep == 1 {
		fmt.Printf("|---%s\n", filepath.Base(dirPath))
	}

	//windows下的目录分割符号是\
	//linux下的目录分隔符是/
	sep := string(os.PathSeparator)
	for _, fi := range dir {
		//
		if fi.IsDir() {
			fmt.Printf("|")
			for i := 0; i<deep; i++{
				fmt.Printf("    |")
			}
			fmt.Printf("----%s\n", fi.Name())
			ListDir(dirPath+sep+fi.Name(), deep+1)
			continue
		}
		fmt.Printf("|")
		for i := 0; i<deep; i++ {
			fmt.Printf("    |")
		}
		fmt.Printf("----%s\n", fi.Name())
	}
	return nil
}

func main(){
	app := cli.NewApp()
	app.Name = "tree"
	app.Usage = "list the dir."

	app.Action = func(c *cli.Context) error{
		var dir string = "."
		if c.NArg() > 0 {
			dir = c.Args()[0]
		}
		ListDir(dir, 1)
		return nil
	}
	app.Run(os.Args)
}
