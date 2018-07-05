package main

import "fmt"

func main() {

	var str string = "hello1"
	fmt.Println(string(str[len(str)-1]))

	linkSet2IPCmd := fmt.Sprintf("ip netns exec %s ip addr add 10.0.%s%s.%s/24 dev %s",
		"a",
		"b",
		"c",
		"d")
}