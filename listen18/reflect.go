package main

/*
1.变量的内在机制
	A.类型信息，这部分是元信息，是预先定义好的
	B.值类型，这部分是程序运行过程中，动态改变的
	反射就是这两部分信息的获取

2.反射与空接口
	a.空接口可以存储任意类型
	b.那么给你一个空接口，怎么里面存储的东西是什么？
	c.在运行时动态的获取一个变量的类型信息和之信息，就叫反射

3.怎么分析？
	a.内置包 reflect
	b.获取类型信息： reflect.TypeOf
	c.获取值信息： reflect.ValueOf

4.
*/

import (
	"fmt"
	"reflect"
)

func reflect_example(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("type of a is: %v\n", t)

	k := t.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64")
	case reflect.String:
		fmt.Printf("a is string")
	}
}

func reflect_value(a interface{}){
	v := reflect.ValueOf(a)
	t := v.Type()   //t := reflect.TypeOf(a)
	fmt.Printf("type of a is: %v\n", t)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("a is int64, store value is :%d", v.Int())
	case reflect.Float64:
		fmt.Printf("a is float64, store value is: %f", v.Float())
	}

}


func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)  //v拿到的是一个副本，不能直接修改值
	k := v.Kind()
	switch k {
	case reflect.Int64:
		v.SetInt(100)
		fmt.Printf("a is int64, store value is :%d", v.Int())
	case reflect.Float64:
		v.SetFloat(6.8)
		fmt.Printf("a is float64, store value is: %f", v.Float())
	case reflect.Ptr:   //x 是值类型，传进来的是副本，所以需要传指针进来修改
		v.Elem().SetFloat(6.8)   //v.Elem()取得内存
		//fmt.Printf("a is float64, store value is: %f", v.Float())
	}
}

func main() {
	var x float64 = 3.4

	//reflect_example(x)
	//reflect_value(x)
	reflect_set_value(&x)
	fmt.Printf("x store value is :%v\n", x)


}
