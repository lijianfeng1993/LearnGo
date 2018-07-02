package main

import (
	"fmt"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
	)

func calcoCoin(username string) int {
	var sum int = 0
	for _, char := range username {
		switch char {
		case 'a','A' :
			sum = sum + 1

		case 'e','E' :
			sum = sum + 1

		case 'i','I' :
			sum = sum +2

		case 'o', 'O' :
			sum = sum + 3

		case 'u','U' :
			sum = sum +5
		}
	}
	return sum
}
func distri(){
	for _, username := range users{
		allCoin := calcoCoin(username)
		distribution[username] = allCoin
	}
	fmt.Println(distribution)
}



func main(){
	distri()
}