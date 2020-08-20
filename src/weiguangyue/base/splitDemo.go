package main

import "fmt"

func display(arr []int) {
	for i := 0; i < len(arr); i++ {

		fmt.Println(arr[i])
	}
}

func foo(arr []int) {

}

func returnSlice() []int {

	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, i)
	}
	return arr
}

func main() {

	{
		var arr = make([]int, 3, 5)

		for i := 0; i < len(arr); i++ {

			arr[i] = i
		}
		//	display(arr)
		foo(arr)
	}

	fmt.Println("............")

	{
		var arr = []int{}

		for i := 0; i < 10; i++ {

			arr = append(arr, i)
		}

		//	display(arr)
		foo(arr)
	}

	fmt.Println("............")

	{
		arr := returnSlice()
		//	display(arr)
		foo(arr)
	}

	{
		arr := returnSlice()

		arr1 := arr[0:2]
		fmt.Println(arr1)

		arr2 := arr[2:3]
		fmt.Println(arr2)

	}

}
