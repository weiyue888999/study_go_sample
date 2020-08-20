package main

import (
	"fmt"
)

func f(ch chan int, x int) {

	fmt.Println(x)

	ch <- x
}

func main() {

	//chh := make(chan interface{})

	count := 10

	ch := make(chan int)
	fmt.Println(ch)

	for i := 0; i < count; i++ {
		go f(ch, i)
	}

	for i := 0; i < count; i++ {
		x := <-ch
		fmt.Println(x)
	}
}
