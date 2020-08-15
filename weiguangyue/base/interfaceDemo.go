package main

import "fmt"

type Service interface {
	start()
	stop()
}

type AppServer struct {
	name string
	port int
}

func (s *AppServer) start() {

	fmt.Println("start")
}

func (s *AppServer) stop() {

	fmt.Println("stop")
}

func test_servie(s Service) {

	s.start()
	s.stop()
}

func main() {

	//在栈内创建结构体变量
	var appServer = AppServer{name: "app", port: 9090}
	appServer.start()
	appServer.stop()

	//在堆内存内创建结构体变量
	var app = new(AppServer)
	app.start()
	app.stop()
	//
	//	var s* Service = AppServer{name: "app", port: 9090}
	//
	//	s.start()
	//	s.stop()

	var app1 = AppServer{name: "app", port: 9090}
	test_servie(&app1)
}
