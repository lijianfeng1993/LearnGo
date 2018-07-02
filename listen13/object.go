package main

import (
	"fmt"
	"encoding/json"
)

//面向对象
//函数定义
//在函数面前加上一个接受着，这样编译器就知道这个方法属于哪个类型了
type People struct {
	Name string
	Country string
}

func (p People) Print(){
	fmt.Printf("Name: %s, Country: %s\n", p.Name, p.Country)
}
func (p People) Set(name, country string) {
	p.Name = name
	p.Country = country
}
func test1() {
	var p1 People = People{
		Name:"ljf",
		Country:"china",
	}
	p1.Print()
	p1.Set("ljf2", "english")   //值类型，在方法中拷贝一份，所以p1平不会被修改
	p1.Print()
}

//go无法为不是包里的结构定义方法

//指针类型作为接受者的区别
func (p *People) Set1(name, country string){
	p.Name = name
	p.Country = country
}
func test2(){
	var p1 People = People{
		Name:"ljf",
		Country:"china",
	}
	p1.Print()
	(&p1).Set1("ljf2", "english") //指针类型，可以修改P1的值
	p1.Print()
}

//什么时候使用值类型/指针类型作为接受者
//	A.需要修改接受者中的值的时候用指针类型
//	B.接受者是大对象的时候，拷贝副本代价比较大
//	C.一般来说，通常使用指针类型作为接受者

//匿名结构体和继承
type Animal struct {
	Name string
	Sex string
}
func (a *Animal)  Talk(){
	fmt.Printf("i am talk, i am %s\n", a.Name)
}

type Dog struct {
	Feet string
	*Animal   //继承于Animal, 通过匿名字段组合的方式实现继承的
}
func (d *Dog) Eat(){
	fmt.Printf("dog is eat.")
}
func (d *Dog) Talk(){
	fmt.Println("i am dog")
}

func test3(){
	var d *Dog = &Dog{
		Feet:"four feet",
		Animal: &Animal{
			Name:"bxh",
			Sex:"male",
		},
	}
	d.Eat()
	d.Talk() //当前实例有就调用当前实例的方法，没有就调用匿名字段中的方法，即父类中的方法
	d.Animal.Talk()  //调用父类中的方法
}

//json序列化(结构体--->json)
type Student struct {
	Id string
	Name string
	Sex string
}
type Class struct {
	Name string
	Count int
	Student []*Student
}
func test4(){
	c := &Class{
		Name:"01",
		Count:0,
	}
	for i:=0; i<10; i++ {
		stu := &Student{
			Name:fmt.Sprintf("stu%d",i),
			Sex:"male",
			Id: fmt.Sprintf("%d",i),
		}
		c.Student = append(c.Student, stu)
	}
	data, err := json.Marshal(c)  //将结构体序列化为json格式
	if err != nil {
		fmt.Println("json marshal failed.")
		return
	}
	fmt.Printf("json: %s\n", string(data))
	/*
	json: {"Name":"01","Count":0,"Student":[{"Id":"0","Name":"stu0","Sex":"male"},{"Id":"1","Name":"stu1","Sex":"male"},{"Id":"2","Name":"stu2","Sex":"male"},{"Id":"3","Name":"stu3","Sex":"male"},{"Id":"4","Name":"stu4","Sex":"male"},{"Id":"5","Name":"stu5","Sex":"male"},{"Id":"6","Name":"stu6","Sex":"male"},{"Id":"7","Name":"stu7","Sex":"male"},{"Id":"8","Name":"stu8","Sex":"male"},{"Id":"9","Name":"stu9","Sex":"male"}]}
	 */
}

//json反序列化
//json写法 ： {“student”:[{"name":"ljf","sex":"male"},{"name":"zh","sex":"male}]}
//json用反引号写
func test5(){
	rawJson := `{"Name":"01","Count":0,"Student":[{"Id":"0","Name":"stu0","Sex":"male"},{"Id":"1","Name":"stu1","Sex":"male"},{"Id":"2","Name":"stu2","Sex":"male"},{"Id":"3","Name":"stu3","Sex":"male"},{"Id":"4","Name":"stu4","Sex":"male"},{"Id":"5","Name":"stu5","Sex":"male"},{"Id":"6","Name":"stu6","Sex":"male"},{"Id":"7","Name":"stu7","Sex":"male"},{"Id":"8","Name":"stu8","Sex":"male"},{"Id":"9","Name":"stu9","Sex":"male"}]}`
	fmt.Println(rawJson)
	var c1 *Class = &Class{}
	err := json.Unmarshal([]byte(rawJson), c1)
	if err != nil {
		fmt.Println("unmarshal failed.")
		return
	}
	fmt.Printf("c1: %#v\n", c1)   //其中学生信息切片中存的是所有学生的地址
	for _, v := range c1.Student {  //打印所有的学生信息
		fmt.Printf("student: %#v\n",v)
	}
}

func main(){
	test5()
}

