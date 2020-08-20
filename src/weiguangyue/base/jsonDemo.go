package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func test1() {

	s := &Student{
		Name: "weiguangyue",
		Age:  30,
	}

	str, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("test encoding....")
		fmt.Println(str)
		fmt.Println(string(str))

		fmt.Println("test decoding...")
		jsonStr := string(str)

		var decodingStu Student
		err := json.Unmarshal([]byte(jsonStr), &decodingStu)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(decodingStu)
		}
	}
}

func test2() {

	{
		//正确的json字符串
		fmt.Println("..................")
		var str string = `{"name","weigungyuae","age":19}`
		var stu Student
		err := json.Unmarshal([]byte(str), &stu)
		if err != nil {

			fmt.Println(err)
		} else {

			fmt.Println(stu)
		}
	}
	{
		//错误的json字符串
		fmt.Println("..................")
		var str string = `{"name":weigungyuae","age":19}`
		var stu Student
		err := json.Unmarshal([]byte(str), &stu)
		if err != nil {

			fmt.Println(err)
		} else {

			fmt.Println(stu)
		}

	}
	{
		//json字符串有的字段，而结构体内没有，则这样的字段不会填充到字段当中
		//当然，这样的字段也无法填充到结构体当中！！！
		fmt.Println("..................")
		var str string = `{"name":"weigungyuae","age":19,"sex":"man"}`
		var stu Student
		err := json.Unmarshal([]byte(str), &stu)
		if err != nil {

			fmt.Println(err)
		} else {

			fmt.Println(stu)
		}

	}

}

func main() {
	test1()
	test2()
}
