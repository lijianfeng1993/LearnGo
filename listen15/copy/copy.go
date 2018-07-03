package main

import (
	"fmt"
	"os"
	"io"
)

func CopyFile(dstFile string, srcFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		fmt.Println("open file error.")
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstFile, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open file error.")
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
	
}
func main(){
	_, err := CopyFile("b.txt", "c.txt")
	if err != nil {
		fmt.Println("copy file failed, err:", err)
	}
	fmt.Println("copy done.")
}