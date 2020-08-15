package main

import "fmt"

type man struct{
	name string
	sex string
}

type student struct{
	man struct
	id string
}

func main() {
	fmt.Println("Hello,world")
}
