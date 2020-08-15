package main

import "fmt"

type student struct {
	name string
	age  int
}

func main() {

	//简单数据类型
	{
		var x int = 10
		y := 100

		fmt.Println(x)
		fmt.Println(y)
	}

	//浮点数也是简单类型
	{
		var x float32 = 1.1
		var y float64 = 2.2

		fmt.Println(x)
		fmt.Println(y)

	}
	//golang竟然支持块级作用域
	{
		var name = "weiguangyue"
		sex := "man"

		fmt.Println(name)
		fmt.Println(sex)
	}

	//复杂类型的变量初始化就复杂了
	//复杂类型的变量，本身也要参数等号右侧了!!!
	{

		var s student = student{

			name: "weiguangyue",
			age:  29,
		}
		fmt.Println(s)
	}

	//数组
	{
		var names [3]string = [3]string{"wei", "guangyue", "yue"}
		fmt.Println(names)

		sex1 := [2]string{"man", "woman"}
		sex2 := [...]string{"man", "woman"}
		fmt.Println(sex1)
		fmt.Println(sex2)
	}

	//结构体可直接声明
	{
		type xx struct {
			name string
			sex  string
		}
		var x_instance xx = xx{name: "wei", sex: "xx"}
		fmt.Println(x_instance)
	}

	fmt.Println("Hello,world")
}
