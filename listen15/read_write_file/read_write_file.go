package main

//文件分类：
//	文本文件和二进制文件
//	二进制文件省空间
//文件存取方式：随机存取和顺序存取

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"io/ioutil"
)

func open_file() {
	// 只读方式打开文件
	file, err := os.Open("./file.go")
	if err != nil {
		fmt.Println("open file failed.")
		return
	}
	defer file.Close()
}

// 文件读取，文件内容很少
func read_file() {
	// 文件读取
	inputfile, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println("open file failed.")
		return
	}
	defer inputfile.Close()
	var buf [128]byte
	n, err := inputfile.Read(buf[:])
	fmt.Println("length:", n)
	fmt.Println("data:", string(buf[0:n]))
}

//读取完整文件
func read_whole_file() {
	inputfile, err := os.Open("read_write_file.go")
	if err != nil {
		fmt.Println("open file fialed.")
	}
	defer inputfile.Close()

	var buf [128]byte
	var content []byte
	for {
		n, err := inputfile.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file error.")
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

//使用bufio读取文件，读取效率更高，但出于应用层的，如果程序挂了，会造成缓存丢失，使用时，视情况而定
func read_bufio_file() {
	file, err := os.Open("read_write_file.go")
	if err != nil {
		fmt.Println("open file failed.")
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("read file error:", err)
			return
		}
		fmt.Println(line)
	}
}

// 通过io/ioutil 工具读取整个文件，简化逻辑处理，非常简单
func ioutil_read_file(){
	content, err := ioutil.ReadFile("read_write_file.go")
	if err != nil {
		fmt.Println("read file error:", err)
		return
	}
	fmt.Println(string(content))
}

// 文件写入 os.OpenFile("file", 第二参数， 第三参数)
// 写入模式, 第二参数：
//		1. os.O_WRONLY    只写
//		2. os.O_CREATE    创建文件
//		3. os.O_ORDONLY	  只读
//		4. os.O_RDWR      读写
//		5. os.O_TRUNC     清空文件
//		6. os.O_APPEND    追加文件
// 权限控制， 第三参数  例子：0666，所有用户都可读和写
//		r --> --> 004
//		w --> --> 002
//		x --> --> 001
func write_file(){

}


func main() {
	//read_whole_file()
	//read_bufio_file()
	ioutil_read_file()
}
