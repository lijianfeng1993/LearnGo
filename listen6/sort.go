package main

import(
	"fmt"
)

//插入排序
func insert_sort(a [8]int) [8]int{
	for i:=1; i< len(a); i++ {
		for j:=i; j>0; j-- {
			if a[j] < a[j-1] {
				a[j],a[j-1] ; a[j-1],a[j]
			} else {
				break
			}
		}
	}
	return a
}



func main(){
	var i [8]int = [8]int(8,3,2,5,12,9,34,66)
	insert_sort(i)
	fmt.Println(i)
}