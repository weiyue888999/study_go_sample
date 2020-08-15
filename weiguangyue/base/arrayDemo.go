package main

import "fmt"
import "os"
import "strconv"

func displayArray2(p *[2]string) {

	for i := 0; i < len(*p); i++ {
		fmt.Println((*p)[i])
	}
}

func displayArray(names [2]string) {

	fmt.Printf("%p", &names)
	fmt.Println()

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}

//这里没有指定数组参数
func displayIntArray(ns []int) {

	for i := 0; i < len(ns); i++ {
		fmt.Println(ns[i])
	}
}

func main() {
	{

		//只能使用切片来创建动态的连续数据
		if len(os.Args) > 1 {

			fmt.Println("create one array")

			n := os.Args[1]
			fmt.Println(n)
			nn, err := strconv.ParseInt(n, 10, 32)
			if err == nil {

				xarr := make([]int, nn)
				fmt.Println(xarr)
			} else {
				fmt.Println(err.Error)
				fmt.Println("wrong number")
			}
		}
	}

	{

		//完整声明形式
		var classes [2]string = [2]string{"1", "2"}
		for i, v := range classes {
			fmt.Println(i)
			fmt.Println(v)
			fmt.Println("...")
		}
		fmt.Println(classes)
	}

	{

		//自动推导的形式
		sexs := [2]string{"man", "woman"}
		fmt.Println(sexs)
	}

	{

		//自动推导长度的形式
		var names = [...]string{"wei", "haha"}
		fmt.Println(names)

		fmt.Printf("%p", &names)
		fmt.Println()

		displayArray(names)

		var p *[len(names)]string = &names
		displayArray2(p)

	}

	{

		arr := []int{1, 9: 200}
		displayIntArray(arr)
		fmt.Println("......")
	}
	//
	//	{
	//
	//		arr := []int{9, 200}
	//		var xarr [2]int = arr
	//		fmt.Println(xarr)
	//		displayIntArray(xarr)
	//	}
	fmt.Println("Hello,world")
}
