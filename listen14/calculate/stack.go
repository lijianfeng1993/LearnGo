package calculate

// 实现一个简易的计算器，支持加减乘除以及带括号的计算表达式，用户从终端输入表达式，程序输出计算结果

//计算表达式通常为中缀表达式，装换成计算机中的后缀表达式
// 9+(3-1)*3+10/2    -->  9 3 1 - 3 * 10 2 / + +

//用到堆栈，用数组实现

import (
	"fmt"
)

type stack struct {
	data [1024]string
	top  int
}

//push data to stack
func (s *stack) Push(d string) {
	s.data[s.top] = d
	s.top ++
}

//pop data from stack
func (s *stack) Pop() (ret string, err error) {
	if s.top == 0 {
		fmt.Println("stack is empty.")
	}
	s.top --
	ret = s.data[s.top]
	return
}

//get top data of stack
func (s *stack) Top() (ret string, err error) {
	if s.top === 0{
		fmt.Println("stack is empty.")
	}
	ret = s.data[s.top--]
	return
}

func (s *stack) Empty() (b bool) {

}