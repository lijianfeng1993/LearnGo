package main

import (
	"fmt"
	"github.com/lijianfeng1993/kubernetes/pkg/util/json"
)

//Go中面向对象是通过struct实现的，struct使用户自定义的类型

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

//结构体初始化
func test1() {
	var user User
	user.Age = 18
	user.AvatarUrl = "http://www.baidu.com"
	user.Username = "ljf"
	user.Sex = "male"

	fmt.Printf("user.username = %s ,age = %d, sex = %s, avatarurl = %s\n", user.Username, user.Age,
		user.Sex, user.AvatarUrl)

	//另一种初始化方法
	user2 := User{
		Username:  "zhanghao",
		Age:       18,
		Sex:       "mail",
		AvatarUrl: "http://www.baodu.com",
	}
	fmt.Printf("user2=%#v\n", user2)
}

//结构体类型的指针
//  var user *User = &User{}
//  fmt.Printf("%p  %#v\n", user)
//&User{} 和 new(User)本质上是一样的，都是返回一个结构体的地址
func test2() {
	//方法1
	var user *User = &User{
		Username:  "zhanghao",
		Age:       18,
		Sex:       "mail",
		AvatarUrl: "http://www.baodu.com",
	}
	//方法2
	var user2 *User = &User{}
	user2.Age = 18 //(*user2).Age = 18 简化之前的写法，语法糖
	user2.AvatarUrl = "http://www.baidu.com"
	user2.Username = "ljf"
	user2.Sex = "male"

	fmt.Println(user2)
	fmt.Println(user)

}

//结构体的内存布局：占用一段连续的内存空间

//结构体没有构造函数，必要时需要自己实现，见user/user.go, usertest.go
//可通过构造函数生成一个对象，不需要再自己定义一个对象

//匿名字段
//匿名字段默认采用类型名作为字段名
type User2 struct {
	Username string
	Sex      string
	int      //字段没有名字
	string
}

func test3() {
	var user User2
	user.Username = "ljf"
	user.Sex = "male"
	user.int = 18
	user.string = "www.baidu.com"
}

//结构体嵌套,一个结构体里面嵌套另一个结构体
type Address struct {
	Province string
	City     string
}
type User3 struct {
	Username string
	Sex      string
	address  Address
}

func test4() {
	user := &User3{
		Username: "ljf",
		Sex:      "male",
		address: Address{
			Province: "beijing",
			City:     "beijing",
		},
	}
	fmt.Printf("user=%#v\n", user)

	var user2 User3
	user2.Username = "zhanghao"
	user2.Sex = "nan"
	user2.address = Address{
		Province: "bk",
		City:     "bk",
	}
	//user2.Province = "bk"   会往里自动找到嵌套结构体中的字段
	//user2.City = "bk"
}

//匿名字段冲突解决,都有city字段,Createtime字段
type Address2 struct {
	Province   string
	City       string
	Createtime string
}
type Email struct {
	Createtime string
}
type User4 struct {
	Username string
	Sex      string
	City     string
	*Address2
	*Email
}
func test5() {
	var user User4
	user.Username = "ljf"
	user.Sex = "nan"
	user.City = "beijing" //优先使用结构体中的字段，

	user.Address2 = new(Address2)
	user.Address2.City = "beijing01"
	fmt.Printf("user=%#v\n", user)
	fmt.Println(user.Address2.City)

	user.Email = new(Email)

	//user.Createtime  = "111" 报错，冲突
	//需要指定那个匿名字段中的Createtime
	user.Address2.Createtime = "111"
	user.Email.Createtime = "222"

	fmt.Printf("user=%#v\n", user)
	fmt.Println(user.Address2.City, user.Email.Createtime)
}


//字段可见性，大写表示外部可访问，小写表示外部不可访问
type User5 struct {
	Username string
	sex string   //外部不可访问
}

//tag是结构体的元信息，可以在运行的时候通过反射的机制读取出来
//字段类型后面，以反引号括起来的key-value结构的字符串，多个tag以逗号隔开
type User6 struct {
	Username string `json:"username"`   //在json中，Username会变成username
	Sex string `json:"sex"`
	Score float32
	age int
}
func test6(){
	user := &User6 {
		Username: "ljf",
		Sex: "nan",
		Score:99,
		age:18,
	}
	data, _ := json.Marshal(user)
	fmt.Printf("json str:%s\n", string(data))  //json str:{"username":"ljf","sex":"nan","Score":99}
	//json序列化后里面没有age，因为小写的访问不了
}

//课后左右，实现一个学生管理系统，每个学生有分数，年级，性别，名字等字段，
//用户可以在控制台添加学生，修改学生信息，打印所有学生列表功能
//代码student

func main() {
	test6()
}
