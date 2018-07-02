package main
//使用面向对象的方式实现学生管理系统
//实现一个学生管理系统，每个学生有分数，年级，性别，名字等字段，用户可以在控制台添加学生，修改学生信息，打印所有学生列表等。

import(
	"fmt"
	"os"
)

type Student struct {
	Username string
	Score string
	Grade string
	Sex string
}

type StudentMgr struct {
	allStudent []*Student
}

func (p *StudentMgr) AddStudent(stu *Student) (err error){
	for index , v := range p.allStudent{
		if v.Username == stu.Username {
			fmt.Println("user %s success update\n", stu.Username)
			p.allStudent[index] = stu
			return
		}
	}
	p.allStudent = append(p.allStudent, stu)
	fmt.Printf("user %s success insert\n", stu.Username)
	return
}

func (p *StudentMgr) ModifyStudent(stu *Student) (err error){
	for index , v := range p.allStudent{
		if v.Username == stu.Username {
			fmt.Println("user %s success update\n", stu.Username)
			p.allStudent[index] = stu
			return nil
		}
	}
	fmt.Printf("user %s is not fount\n",stu.Username)
	err = fmt.Errorf("user %s is not exist.")
	return
}

func (p *StudentMgr) ShowStudent(){
	for _, v := range p.allStudent{
		fmt.Printf("user: %s info: %#v",v.Username, v)
	}
	fmt.Println()
}

var (
	studentMgr = &StudentMgr{}
)

func NewStudent(username string, sex string, grade string, score string) *Student {
	stu := &Student{
		Username: username,
		Sex:      sex,
		Grade:    grade,
		Score:    score,
	}
	return stu
}
func inputStudent()  (*Student){
	var (
		username string
		sex      string
		grade    string
		score    string
	)
	fmt.Printf("Please input username:")
	fmt.Scanf("%s\n", &username)
	fmt.Printf("Please input sex[0/1]:")
	fmt.Scanf("%s\n", &sex)
	fmt.Printf("Please input grade[0-6]:")
	fmt.Scanf("%s\n", &grade)
	fmt.Printf("Please input score:")
	fmt.Scanf("%s\n", &score)
	stu := NewStudent(username, sex, grade, score)
	return stu
}

func main(){
	for {
		//showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			stu := inputStudent()
			studentMgr.AddStudent(stu)
			/*
				case 2:
					ModifyStudent()
			*/
		case 3:
			studentMgr.ShowStudent()
		case 4:
			os.Exit(0)
		}
	}
}