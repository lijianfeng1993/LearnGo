package maim

import (
	"fmt"
)

func testIf() {
	var a int = 2
	if a == 1 {
		fmt.Println("a = 1\n")
	} else if a == 2 {
		fmt.Println("a = 2\n")
	} else if a == 3 {
		fmt.Println("a = 3\n")
	}
}

func testSwitch() {
	var a int = 2
	switch a {
	case 1:
		fmt.Println("a = 1\n")
	case 2:
		fmt.Println("a = 2\n")
	case 3:
		fmt.Println("a = 3\n")
	case 4,5,6:
		fmt.Println("a=4/5/6")
	default:
		fmt.Println("aaaaa")
	}
}

func testSwitch2() {
	var num int = 60
	switch {
	case num >20 && num <70:
		fmt.Println("ok")
	}
}

func main() {

}
