package main

import (
	"fmt"
	"image/color"
)

type ColoredPoint struct {
	color.Color     //嵌入字段
	x, y        int //聚合字段
}

type Count int
type StringMap map[string]string
type FloatChan chan float32
type MyFunc func(int) (int, error)

//接收者是指针才能够改变自身
func (c *Count) Add1(x int) {

	cv := int(*c)
	cv += x
	*c = Count(cv)
}

//接收者是值，则不能改变自身
func (c Count) Add2(x int) {
	cv := int(c)
	cv += x
	c = Count(cv)
}

func test_add() {

	fmt.Println("...........")

	var i Count = 100
	fmt.Println("count:", i)
	i.Add1(1)

	fmt.Println("count:", i)

	i.Add2(100)

	fmt.Println("count:", i)

}

func test_color() {

	fmt.Println("..............")

	cp := ColoredPoint{}
	fmt.Println(cp)

	fmt.Println(cp.Color)
	fmt.Println(cp.x)
	fmt.Println(cp.y)
}

func test_type() {

	fmt.Println("..............")

	var i Count = 1
	i++
	fmt.Println(i)

	var m StringMap = make(StringMap)
	m["wei"] = "weiguangyue"
	fmt.Println(m)

	var fc FloatChan = make(FloatChan)
	fmt.Println(fc)
}

func test_type_func() {

	var myFunc MyFunc = func(x int) (int, error) {

		return x, nil
	}

	result, err := myFunc(100)
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println(result)
	}

}

func main() {

	test_color()
	test_type()
	test_type_func()
	test_add()
}
