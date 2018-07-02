package main

import (
	"fmt"
	"time"
)

func cal() {
	for i := 0; i<10; i++ {
		time.Sleep(1*time.Second)
		fmt.Println("run" + " times")
	}
	fmt.Println("cal finish.")
}

func main() {
	go cal()
	fmt.Println("main finish.")
	time.Sleep(time.Second * 11)
}