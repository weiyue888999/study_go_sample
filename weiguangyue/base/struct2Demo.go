package main

import "fmt"

type student struct {
	name string
}

func (s *student) SetName(name string) {

	this := *s
	this.name = name
}

func (s student) GetName() string {

	return s.name
}

func test(s *student) {
	fmt.Println(s)
}

func main() {

	var s student = student{name: "wei"}
	fmt.Println(s.GetName())

	s.SetName("new name")
	fmt.Println(s.GetName())

	fmt.Println("............")

	ps := &s
	var str string = ps.GetName()
	fmt.Println(str)

	test(&s)
	test(ps)

	i := 10
}
