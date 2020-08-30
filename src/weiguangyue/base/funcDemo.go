package main

import (
	"fmt"
)

type Student struct {
	name string
}

func f1(s *Student) {

	fmt.Println(s)
	fmt.Println("f1")
}

func f2(s Student) {

	fmt.Println(s)
	fmt.Println("f2")
}

//普通函数，声明是指针，就要传递指针，声明是值，就要传值
//和结构体的方法的接收者是不一样的
func main() {

	ps := &Student{name: "wei"}

	s := Student{name: "weiguangyue"}

	f1(ps)
	f1(&s)

	f2(*ps)
	f2(s)
}
