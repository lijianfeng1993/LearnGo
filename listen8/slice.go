package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//切片是基于数组类型做的一层封装，非常灵活，可以自动扩容,切片是对数组的引用
//var a []int  没有写长度就是切片，写了长度就是数组
func testSlice0() {
	var a []int //空切片
	if a == nil {
		fmt.Printf("a is nil\n")
	} else {
		fmt.Printf("a = %v\n", a)
	}
	//a[0] = 100, 报错，没有初始化
}

//切片初始化，a[start:end]创建一个包括从start到end-1的切片
func testSlice1() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int = a[1:4] //基于数组a创建一个切片，包括元素a[1],a[2],a[3]
	fmt.Println(b)
	fmt.Println(b[0])

}

//切片初始化方法2
func testSlice2() {
	c := []int{6, 7, 8} //创建一个数组并返回一个切片，中括号里没有长度
	fmt.Println(c)
}

//切片的基本操作
//arr[start:end] 包括start到end-1(包括end-1)之间的所有元素
//arr[start:]
//arr[:end]
//arr[:]  包括整个数组的所有元素

//切片修改操作
func testSlice3() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a)
	b := a[1:3]
	b[0] = b[0] + 10
	b[1] = b[1] + 20
	fmt.Println(b)
	fmt.Println(a) //{1,12,23,4,5}
	/*
		for index, val := range b{
			b[index] = b[index] + 10
		}
	*/
}

//使用make创建切片,扩容操作
func testMake1() {
	var a []int
	a = make([]int, 5, 10) //切片长度5，切片容量10，底层是一个长度为10的数组
	a[0] = 10
	a[1] = 20
	fmt.Printf("a = %v， len = %d, cap = %d\n\n", a, len(a), cap(a)) // a = [10,20,0,0,0]

	a = append(a, 11)
	fmt.Printf("a = %v， len = %d, cap = %d\n\n", a, len(a), cap(a)) // a = [10,20,0,0,0,11]

	b := make([]int, 0, 10)
	b = append(b, 10)

	//当append的数量大于容量的时候，就会自动扩容，底层重新生成一个更大的数组，并基于新的数组生成切片。容量翻倍，并把前面的值拷贝进去。
}

//切片容量
func testCap() {
	a := [...]string{"a", "b", "c", "d", "e", "f"}
	b := a[1:3]
	fmt.Printf("b=%v  len=%d  cap=%d\n", b, len(b), cap(b)) //容量为5，因为b从1开始，容量是从切片开始的位置开始算
}

//切片再切片
func testCap2() {
	a := [...]string{"a", "b", "c", "d", "e", "f"}
	b := a[1:3]
	fmt.Printf("b=%v  len=%d  cap=%d\n", b, len(b), cap(b)) //b=[b c]  len=2  cap=5

	b = b[:cap(b)]
	fmt.Printf("b=%v  len=%d  cap=%d\n", b, len(b), cap(b)) //b=[b c d e f]  len=5  cap=5
}

//空切片
func testSlice4() {
	var a []int
	//a[0] = 100  会报错
	if a == nil {
		fmt.Printf("a is nil.")
	}

	fmt.Printf("%p, len=%d, cap=%d\n", a, len(a), cap(a)) //a is nil.0x0, len=0, cap=0

	a = append(a, 100)
	fmt.Printf("%p, len=%d, cap=%d\n", a, len(a), cap(a)) //0xc42007a008, len=1, cap=1

	a = append(a, 200)
	fmt.Printf("%p, len=%d, cap=%d\n", a, len(a), cap(a)) //0xc42007a030, len=2, cap=2

	//翻倍扩容
	a = append(a, 300)
	fmt.Printf("%p, len=%d, cap=%d\n", a, len(a), cap(a)) //0xc420090020, len=3, cap=4

	a = append(a, 400)
	fmt.Printf("%p, len=%d, cap=%d\n", a, len(a), cap(a)) //0xc420090020, len=4, cap=4
}

//追加一个切片到另一个切片的末尾
//切片后面的三个点表示展开切片成一个个元素
func testSlice5() {
	a := []int{1, 3, 4}
	b := []int{2, 5}
	a = append(a, b...)
	fmt.Println(a)
}

//切片传参，函数的参数为切片
//求切片中所有元素之和
func sum(a []int) int {
	var sum int = 0
	for _, val := range a {
		sum += val
	}
	return sum
}
func testSum() {
	a := []int{1, 2, 3, 4, 5}
	s := sum(a)
	fmt.Println(s)
}

//修改,因为切片是引用，所以可以修改，区别于数组，看数组的例子（array.go 85行）
func modifySlice(a []int) {
	a[0] = 1000
}
func testModifySlice() {
	a := [5]int{1, 2, 3, 4, 5}
	modifySlice(a[:])
	fmt.Println(a) //[1000 2 3 4 5]
}

//切片的拷贝,不会自动扩容，长度多少就拷贝多少
func testCopy() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := []int{1}
	copy(a, b) //把b拷贝到a里面
	copy(c, b)
	fmt.Println(a) //[4,5,6]
	fmt.Println(b) //[4,5,6]
	fmt.Println(c) //[4]
}

//切片遍历和数组遍历一样

//make 和 new 的区别
//make为内建类型slice、map、channel分配内存
//new用于为各种类型的内存分配，new返回一个指针

//练习1：使用golang标准包“sort”对数组进行排序
func testSort() {
	var a [5]int = [5]int{5, 4, 3, 2, 1}
	sort.Ints(a[:]) //Ints对整数进行排序
	fmt.Println("a: ", a)

	b := [5]string{"ac", "ec", "be", "fa", "ii"}
	sort.Strings(b[:]) //Strings对字符串进行排序
	fmt.Println("b: ", b)

	//sort.Float64s 对浮点数进行排序

}

//练习2：实现一个密码生成工具，支持以下功能：
//  (1)用户可以通过-l指定生成密码的长度
//  (2)用户可以通过-t指定生成密码的字符集，比如-t num生成全数字的密码，-t char生成包含全英文字符的密码，
// 		-t mix包含生成数字和英文的密码，-t advance生成包含数字、英文以及特殊字符的密码
var (
	length int
	charse string
)

const (
	NumStr  = "0123456789"
	CharStr = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func parseArgs() {
	flag.IntVar(&length, "l", 16, "-l 生成密码的长度")
	flag.StringVar(&charse, "t", "num",
		`-t 制定密码生成的字符集，
				num:只使用数字[0-9],
				char:只是用英文字母[a-zA-Z],
				mix:使用数字和字母，
				advance:使用数字、字母以及特殊字符`)
	flag.Parse()
}
func generagePassword() string{
	var passwd []byte = make([]byte, length, length)
	var sourceStr string

	if charse == "num" {
		sourceStr = NumStr
	} else if charse == "char" {
		sourceStr = CharStr
	} else if charse == "mix" {
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr)
	} else if charse == "advance" {
		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr)
	} else {
		sourceStr = NumStr
	}
	//fmt.Println("sourceStr: ", sourceStr)

	for i:=0; i<length; i++ {
		index := rand.Intn(len(sourceStr))
		passwd[i] = sourceStr[index]
	}
	return string(passwd)
}
func testCommand() {
	rand.Seed(time.Now().UnixNano()) //用当前时间生成随机化种子
	parseArgs()
	fmt.Printf("length:%d, charset:%s\n", length, charse)
	passwd := generagePassword()
	fmt.Println(passwd)
}



func main() {
	//testSlice5()
	//testMake1()
	//testCap2()
	//testSum()
	//testModifySlice()
	//testCopy()
	//testSort()
	testCommand()
}
