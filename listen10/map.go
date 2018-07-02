package main

import (
	"fmt"
	"sort"
	"strings"
)

//map类型是一个key-value的数据结构
//声明和定义
//var a map[string]int    key:是string, value是int类型
//var b map[int]string
//var c map[float32]string
//map 必须初始化才能使用，否则panic
//map是一个引用类型，和切片一样
func testmap1() {
	var a map[string]int //声明，但是无法使用，没有初始化
	if a == nil {
		a = make(map[string]int, 16) //容量为16，可以不写容量,Map的扩容会消耗资源，所以尽量预估一下容量
	}
	a["stu01"] = 100
	fmt.Println(a)

	//另一种初始化方法
	//var a map[string]int = map[string]int {
	//   "st01":1000,
	//   "st02":2000,}
}

//map基本操作
//通过key访问元素
func testmap2() {
	a := map[string]int{
		"ljf": 1000,
		"bxh": 2000,
		"zh":  3000,
	}
	var key string = "ljf"
	fmt.Println(a[key])
}

//判断map中key是否存在
func testmap3() {
	a := make(map[string]int)
	a["ljf"] = 1000
	a["zh"] = 2000
	a["bxh"] = 3000
	var result1 int
	var ok1 bool
	result1, ok1 = a["qcl"]
	if ok1 == false {
		fmt.Printf("result1:%d\n", result1)
	} else {
		fmt.Printf("result1:%d\n", result1)
	}
}

//遍历和删除
func testmap4() {
	a := make(map[string]int)
	a["ljf"] = 1000
	a["zh"] = 2000
	a["bxh"] = 3000
	for key, val := range a {
		fmt.Printf("key:%s,value:%d\n", key, val)
	}

	//删除记录
	delete(a, "bxh")
	for key, val := range a {
		fmt.Printf("key:%s,value:%d\n", key, val)
	}

	//删除所有记录
	for key, _ := range a {
		delete(a, key)
	}
	fmt.Println(a)
}

//map是引用类型，里面存的是地址
func testmap5() {
	a := map[string]int{
		"ljf": 1000,
		"bxh": 2000,
		"zh":  3000,
	}
	b := a
	b["ljf"] = 5000
	fmt.Println(a)
}

//默认是无序输出，通过排序好的key进行输出
func testmap6() {
	a := map[string]int{
		"ljf": 1000,
		"bxh": 2000,
		"zh":  3000,
	}
	keys := make([]string, 0, 3)
	for key, _ := range a {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, val := range keys {
		fmt.Printf("key:%s, val:%d\n", val, a[val])
	}

}

//map类型的切片
//切片中可以存任何类型，切片中可以存放map
func testmap7() {
	var s []map[string]int
	s = make([]map[string]int, 5, 16)
	for index, value := range s {
		fmt.Printf("slice[%d]=%v\n", index, value)
	}
	//切片已经初始化，但是切片中的元素没有初始化
	//还需对切片中的元素map进行初始化
	s[0] = make(map[string]int, 16)
	s[0]["ljf"] = 100
	s[0]["zh"] = 200
	s[0]["bxh"] = 300
	for index, value := range s {
		fmt.Printf("slice[%d]=%v\n", index, value)
	}
}

//map中的每个元素是一个切片
func testmap8() {
	var s map[string][]int
	s = make(map[string][]int)
	key := "ljf"
	value, ok := s[key]
	if !ok {
		s[key] = make([]int, 0, 16) //初始化切片
		value = s[key]
	}
	value = append(value, 100)
	value = append(value, 200)
	s[key] = value
	fmt.Println(s)
}

//课后作业1:统计一个字符串中每个单词出现的次数，如s="how do you do",输出how=1,do=2，you:1
func statWordCount(str string) map[string]int {
	var result map[string]int = make(map[string]int, 128)
	words := strings.Split(str, " ")
	for _, val := range words {
		count, ok := result[val]
		if !ok {
			result[val] = 1
		} else {
			result[val] = count + 1
		}
	}
	return result
}
func test1() {
	s := "how do you do"
	result := statWordCount(s)
	fmt.Println(result)
}

//课后作业2：实现学生信息的存储，学生有id，年龄，分数等信息。需要非常方便的通过id查找到对应的学生信息
func studentInfoStore() {
	var stuMap = make(map[int]map[string]interface{}, 100) //空的interface可以存任何类型

	//插入学生id=1,姓名=ljf,分数=78.2,年龄=18
	id := 1
	name := "ljf"
	score := 78.2
	age := 18

	value, ok := stuMap[id]
	if !ok {
		value = make(map[string]interface{}, 8)
		value["name"] = name
		value["id"] = id
		value["score"] = score
		value["age"] = age
	} else {
		value["name"] = name
		value["id"] = id
		value["score"] = score
		value["age"] = age
	}
	stuMap[id] = value
	fmt.Println(stuMap)
}
func test2() {
	studentInfoStore()
}

func main() {
	test2()
}
