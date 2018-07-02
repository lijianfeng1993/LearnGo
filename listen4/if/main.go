package main

import (
	"fmt"
)

func testIf1() {
	num := 10
	if num%2 == 0 {
		fmt.Printf("num:%d is oushu\n", num)
	}else {
		fmt.Printf("num:%d is jishu\n")
	}
}

func testIf2(){
	num := 10
	if num >5 && num <10 {
		fmt.Printf("num:%d hahahah\n", num)
	}else if num >10 && num <20 {
		fmt.Printf("num:%d houhouhou\n", num)
	}else {
		fmt.Printf("num:%d ok\n",num)
	}
}

func testIf3(){
	//num只在if块中有效，出了if就无效了
	if num := 10; num % 2 == 0 {
		fmt.Printf("num:%d is oushu\n",num)
	}else {
		fmt.Printf("num:%d is jishu\n",num)
	}

}

func main(){
	testIf1()
	testIf2()
	testIf3()
}
