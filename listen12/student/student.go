package main

import (
	"fmt"
	"os"
)

type Student struct {
	Username string
	Sex      string
	Grade    string
	Score    string
}

func NewStudent(username string, sex string, grade string, score string) *Student {
	stu := &Student{
		Username: username,
		Sex:      sex,
		Grade:    grade,
		Score:    score,
	}
	return stu
}

var (
	AllStudent []*Student
)

func AddStudent() {
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
	for index, v := range AllStudent {
		if v.Username == stu.Username {
			AllStudent[index] = stu
			return
		}
	}
	AllStudent = append(AllStudent, stu)
}

func ShowStudent() {
	for _, v := range AllStudent {
		fmt.Printf("%#v\n", v)
	}
}

func main() {
	for {
		//showMenu()
		var sel int
		fmt.Scanf("%d\n", &sel)
		switch sel {
		case 1:
			AddStudent()
			/*
				case 2:
					ModifyStudent()
			*/
		case 3:
			ShowStudent()
		case 4:
			os.Exit(0)
		}
	}
}
