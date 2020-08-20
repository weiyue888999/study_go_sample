package main

import "fmt"

func test_string_kv() {

	fmt.Println("map demo")

	m := make(map[string]string, 10)
	m["name"] = "weiguangyue"
	m["age"] = "18"

	name := m["name"]
	fmt.Println(name)

	//当key不存在，则返回0值
	sex := m["sex"]
	fmt.Println(sex)

	//使用这种方式比较稳妥
	sex, found := m["sex"]
	fmt.Println(sex, found)

	//遍历map
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("...")

	//只遍历key
	for k, _ := range m {
		fmt.Println(k)
	}

	fmt.Println("...")

	for _, v := range m {
		fmt.Println(v)
	}

	fmt.Println("...")
}

//老师和学生，一个学生有一个老师

type Student struct {
	name string
}

func (s Student) String() string {

	str := "student[name=" + s.name + "]"
	return str
}

type Teacher struct {
	name string
}

func (t Teacher) String() string {
	str := "teacher[name=" + t.name + "]"
	return str
}

func test_struct_kv() {

	m := make(map[Student]Teacher, 10)
	{

		s := Student{
			name: "魏广跃",
		}

		t := Teacher{
			name: "孔子",
		}

		m[s] = t
	}

	for s, t := range m {

		fmt.Println(s, t)
	}
}

func test_point_kv() {
	m := make(map[*Student]*Teacher, 10)
	{

		s := Student{
			name: "魏广跃",
		}

		t := Teacher{
			name: "孔子",
		}

		m[&s] = &t
	}
	{
		//这样也行！！！
		m[&Student{name: "weiguangyue"}] = &Teacher{name: "kongzi"}

	}

	for s, t := range m {

		fmt.Println(s, t)
	}
}

func test_func_kv() {

	m := make(map[string]func(), 10)
	m["weiguangyue"] = func() {

		fmt.Println("func")
	}
	m["weiguangyue"]()

	//这样会报错
	//m["haha"]()

	f, found := m["haha"]
	if found {

		f()
	}

	//这样写更简短
	if f, found := m["haha"]; found {

		f()
	}
}

func main() {

	//test_string_kv()
	//test_struct_kv()
	//test_point_kv()
	test_func_kv()

}
