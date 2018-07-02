package main

import(
	"fmt"
)

//格式化输入
//从终端获取用户的输入
//  fmt.Scanf(format string, a ...interface{})  格式化输入，空格作为分隔符，占位符和格式化输出一致
//  fmt.Scan(a ...interface{})  从终端获取用户输入，存储在Scanln中的参数里，空格和换行符作为分隔符
//  fmt.Scanln(a ...interface{})  从终端获取用户输入，存储在Scanln中的参数里，空格作为分隔符，遇到换行符结束
func testScanf()  {
	var a int
	var b string
	var c float32
	fmt.Scanf("%d %s %f", &a, &b, &c)	  //传地址进去
	/* 或者
	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%s\n", &s)
	fmt.Scanf("%f\n", &f)
	 */
	fmt.Printf("a=%d, b=%s, c=%f\n", a, b, c)
}
func testScan(){
	var a int
	var b string
	var c float32
	fmt.Scan(&a,&b,&c)
	fmt.Printf("a=%d, b=%s, c=%f\n", a, b, c)
}
func testScanln(){
	var a int
	var b string
	var c float32
	fmt.Scanln(&a,&b,&c)
	fmt.Printf("a=%d, b=%s, c=%f\n", a, b, c)
}

//格式化输入
//从字符串中获取输入
//  fmt.Sscanf(str string, format string, a ...interface{}) 格式化输入，空格作为分隔符，占位符和格式化输出一致
//  fmt.Sscan(str string, a ...interface{})  从终端获取用户输入，存储在Scanln中的参数里，空格和换行符作为分隔符
//  fmt.Sscanln(str string, a ...interface{})  从终端获取用户输入，存储在Scanln中的参数里，空格作为分隔符，遇到换行符结束
func testSscanf(){
	var a int
	var b string
	var c float32
	var str string = "88 adf 8.8"
	// var str string = "88 adf\n\n 8.8"  不可以，换行符不可作为分隔符处理
	fmt.Sscanf(str, "%d%s%f\n", &a, &b, &c)
	fmt.Printf("a=%d, b=%s, c=%f\n", a, b, c)
}
func testSscan(){
	var a int
	var b string
	var c float32
	var str string = "88 adf 8.8"
	// var str string = "88 adf\n\n 8.8"  也可以，换行符也作为分隔符处理
	fmt.Sscan(str, &a,&b,&c)
	fmt.Printf("a=%d, b=%s, c=%f\n", a, b, c)
}
func testSscanln(){
	var a int
	var b string
	var c float32
	var str string = "88 adf 8.8"
	fmt.Sscanln(str, &a,&b,&c)
	fmt.Printf("a=%d, b=%s, c=%f\n", a, b, c)
}

//格式化输出到终端
//fmt.Printf()  格式化输出，并打印到终端
//fmt.Print()  把零个或多个变量打印到终端
//fmt.Println()  把零个或多个变量打印到终端，并换行

//格式化输出到字符串
//fmt.Sprintf(fotmat string, a ...interfaceP{})  格式化并返回字符串
//fmt.Sprint(a ...interface{})  把零个或多个变量按空格进行格式化，返回字符串
//fmt.Sprintln(a ...interface{})  把零个或多个变量按空格进行格式化并换行，返回字符串


func main(){
	testSscanln()

}



