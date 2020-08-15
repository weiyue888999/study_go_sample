package main

import "fmt"

type nameer interface {
	SetName(name string)
	GetName() string
}

type student struct {
	name string
}

func (s *student) String() string {

	name := s.name
	result := "student[name=" +
		name +
		"]"

	return result
}

func (s *student) SetName(name string) {

	s.name = name
}

func (s *student) GetName() string {
	return s.name
}

//----------------

type hightStudent struct {
	student
	sex string
}

func (h *hightStudent) SetSex(sex string) {

	h.sex = sex
}

func (h *hightStudent) GetSex() string {

	return h.sex
}

//----------------

func main() {
	s := student{name: "wei"}
	fmt.Println(s.String())

	s.SetName("guangyue")
	fmt.Println(s.String())

	fmt.Println("................")

	h := hightStudent{sex: "woman"}
	h.name = "weiwei"
	fmt.Println(h)

}
