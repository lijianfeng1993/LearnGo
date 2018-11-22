package main

import (
	"reflect"
	"fmt"
)

type Student struct {
	Name string
	Sex int
	Age int
}

func main(){
	var s Student
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	kind := t.Kind()
	switch kind {
	case reflect.Int64:
		fmt.Println("s is int.")
	case reflect.Float32:
		fmt.Println("s is float.")
	case reflect.Struct:
		fmt.Println("s is struct.")
		fmt.Printf("field num of s is %d\n", v.NumField())
		for i:=0; i<v.NumField(); i++{
			field := v.Field(i)

		}
	default:
		fmt.Println("default.")
	}
}
