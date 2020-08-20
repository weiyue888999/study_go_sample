package main

import "fmt"

func test_return_val(n int) int {

	return n
}

func test_append() {

	arr := make([]int, 10)
	//这里若不适用变量接收，则会报错！！！
	arr = append(arr, 1)

	arr_all := make([]int, 0)
	arr_all = append(arr_all, 2)
	arr_all = append(arr_all, arr...)

	fmt.Println(arr)
	fmt.Println(arr_all)

}

func test_split() {

	arr := make([]int, 0)

	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	fmt.Println(arr)

	//改变底层数组，其余的切片也同时会改变
	{
		arr1 := arr[0:1]
		arr1[0] = 100
		fmt.Println(arr1)
	}

	{

		arr1 := arr[0:2]
		fmt.Println(arr1)
	}
}

func main() {

	test_split()

	test_append()

	//这种普通的函数则不报错！！！
	test_return_val(10)

}
