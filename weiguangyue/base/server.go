package main

import "fmt"
import "strconv"

type service interface {
	start()
	stop()
	status()
}

type server struct {
	port int
	name string
}

func (this *server) start() {
	fmt.Println("start")
}

func (this *server) stop() {

	fmt.Println("stop")
}

func (this *server) String() string {
	return "Server[port:" + strconv.Itoa(this.port) + ",name:" + this.name + "]"
}

func main() {

	s := server{
		port: 200,
		name: "name",
	}

	var ps *server = &s
	ps.start()
	fmt.Println(ps)
	ps.stop()
}
