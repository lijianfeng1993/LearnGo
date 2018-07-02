package main

import (
	"fmt"
	"LearnGo/listen12/user"
)

func main(){
	var user1 user.User
	user1.Age=18
	fmt.Printf("user = %#v\n", user1)

	user2 := user.NewUser("ljf","male",18,"www.baidu.com")
	fmt.Printf("user= %#v\n", user2 )
}

