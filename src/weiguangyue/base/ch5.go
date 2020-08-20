package main

import "fmt"

func test_int(x interface{}) {

	if _, success := x.(int); success {
		fmt.Println("xxx")
	}
}

func test_pass(x int) {

	switch x {
	case 1:
		fmt.Println("1")
		fullthroughg
	case 2:
		fmt.Println("2")
		fullthroughg
	default:
		fmt.Println("not found")
		fullthroughg
	}
}

func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println("Hello,world")
	//	fmt.Println(i)
	//}

	//var name string = "weiguangyue is hero!"

	var i interface{} = 99
	test_int(i)
	test_pass(3)
}
