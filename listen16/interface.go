package main

import (
	"fmt"
)

//接口
//接口定义了一个对象的行为规范
//	只定义规范，不实现
//	具体的对象需要实现规范的细节

//接口的定义
//	type 接口名字 interface
//	接口里面是一组方法签名的集合
/*
	type Animal interface {
		Talk()
		Eat() int
		Run()
	}
 */

//go中接口的实现
//	一个对象中只要包含了接口中的方法，就是实现了这个接口,注意要包含接口中的所有的方法
//	接口类型的变量可以保存具体类型的实例
type Animal interface {
	Talk()
	Eat()
	Name() string
}

type Dog struct {
}
func (d *Dog) Talk(){
	fmt.Println("dog is talking")
}
func (d *Dog) Eat(){
	fmt.Println("dog is eating")
}
func (d *Dog) Name() string{
	fmt.Println("my name is wangcai.")
	return "wangcai"
}

type Pig struct {
}
func (p *Pig) Talk(){
	fmt.Println("pig is talking")
}
func (d *Pig) Eat(){
	fmt.Println("pig is eating")
}
func (d *Pig) Name() string{
	fmt.Println("my name is bajie.")
	return "bajie"
}

func test_interface_animal(){
	var d Dog
	var a Animal
	a = &d  //接口中可以存储任何类型的实现了接口的变量
	a.Eat()
	a.Name()

	var p Pig
	a = &p
	a.Eat()
	a.Name()
}

//一个公司需要计算所有职工的薪水
//每个员工的薪水计算方式不一样
type Employer interface {
	CalcSalary() float32
}

type Programer struct {
	name string
	base float32
	extra float32
}
func NewProgramer(name string, base float32, extra float32) *Programer{
	return &Programer{
		name:name,
		base:base,
		extra:extra,
	}
}
func (p *Programer) CalcSalary() float32{
	return p.base
}

type Sale struct {
	name string
	base float32
	extra float32
}
func NewSale(name string, base float32, extra float32) *Sale{
	return &Sale{
		name:name,
		base:base,
		extra:extra,
	}
}
func (s *Sale) CalcSalary() float32{
	return s.base + s.extra * s.base *0.5
}

func calcAll(e []Employer) float32 {
	var cost float32
	for _, v := range e {
		cost = cost + v.CalcSalary()
	}
	return cost
}
func test_interface_employer(){
	p1 := NewProgramer("ljf", 1500.0, 0)
	p2 := NewProgramer("bxh", 1300.0, 0)
	p3 := NewProgramer("zh", 1300.0, 0)

	s1 := NewSale("a", 800.0, 0.5)
	s2 := NewSale("b", 900.0, 0.8)

	var employerlist []Employer
	employerlist = append(employerlist, p1)
	employerlist = append(employerlist, p2)
	employerlist = append(employerlist, p3)
	employerlist = append(employerlist, s1)
	employerlist = append(employerlist, s2)
	fmt.Printf("这个月总人力成本：%f\n", calcAll(employerlist))
	// 用接口避免了通过各个员工类型计算工资，节省代码，否则，如果有多种员工，需要对每种员工单独计算，再相加
}

//接口类型变量
// 	var a Animal
//	那么a能存储所有实现了Animal接口的对象实例

//空接口
//	空接口没有定义任何方法
//	所以任何类型都实现了空接口
//	空接口可以存储任何类型
func test_interface_null(){
	var a interface{}
	var b int = 10
	a = b
	fmt.Printf("%T %v\n", a, b)

	var c string = "hello"
	a = c
	fmt.Printf("%T %v\n", a, c)

	var d map[string]int = make(map[string]int,100)
	d["a"] = 1
	d["b"] = 2
	a = d
	fmt.Printf("%T %v\n", a, d)
}

func describe(a interface{}) {
	fmt.Printf("%T %v\n", a, a)
}
func test_interface_null2(){
	var b int = 10
	describe(b)
	var c string = "hello"
	describe(c)
	var d map[string]int = make(map[string]int,100)
	d["a"] = 1
	d["b"] = 2
	describe(d)
}

//类型断言
//	如何获取接口类型里面存储的具体类型呢
//	可以把接口类型转成具体的类型
func test(a interface{})  {
	s, ok := a.(int)
	if ok {
		fmt.Println(s)
		return
	}
	str, ok := a.(string)
	if ok {
		fmt.Println(str)
		return
	}
	fmt.Println("can not asser the type of a.")
}
func test_interface_assert(){
	var s int = 10
	test(s)
	var str string = "hello"
	test(str)
	var fl float32 = 1.1
	test(fl)
}
//	通过switch方式断言类型
func testswitch(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("a is int:", a.(int))
	case string:
		fmt.Println("a is string:", a.(string))
	case float32:
		fmt.Println("a is float32:", a.(float32))
	default:
		fmt.Println("not support type.")
	}
	/* 更好的写法
	switch v:=a.(type) {
	case int:
		fmt.Println("a is int:", v)
	case string:
		fmt.Println("a is string:", v)
	case float32:
		fmt.Println("a is float32:", v)
	default:
		fmt.Println("not support type.")
	}
	 */
}
func test_interface_assert2()  {
	var s int = 10
	testswitch(s)
	var str string = "hello"
	testswitch(str)
	var fl float32 = 1.1
	testswitch(fl)
}



func main(){
	//test_interface_animal()
	//test_interface_employer()
	//test_interface_null()
	//test_interface_null2()
	//test_interface_assert()
	test_interface_assert2()
}