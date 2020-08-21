package netbase

import "fmt"

type man struct {
	id int
}

func (m *man) display() {
	fmt.Println(m)
}

type student struct {
	name string
	man
}

/*
func main() {

	m := man{id: 1}

	m.display()

	fmt.Println(".....")

	s := student{
		man: man{
			id: 1,
		},
		name: "wei",
	}
	s.display()
}
**/
