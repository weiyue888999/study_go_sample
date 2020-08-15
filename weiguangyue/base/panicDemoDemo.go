package main

import "fmt"

func main() {
	fmt.Println("Hello,world")

	defer func() {

		fmt.Println("c")
		if err := recover(); err != nil {

			fmt.Println(err)
		}

	}()

	f()
}

func f() {

	fmt.Println("a")
	panic(1111)
	fmt.Println("b")
}
