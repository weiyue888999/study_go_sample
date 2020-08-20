package main

import "fmt"

func test_return_val(n int) int {

	return n
}

func main() {

	fmt.Println("hi")

	arr := make([]int, 10)
	//这里若不适用变量接收，则会报错！！！
	arr = append(arr, 1)

	arr_all := make([]int, 0)
	arr_all = append(arr_all, 2)
	arr_all = append(arr_all, arr...)

	fmt.Println(arr)
	fmt.Println(arr_all)

	//这种普通的函数则不报错！！！
	test_return_val(10)

}
