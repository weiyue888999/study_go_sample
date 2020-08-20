package main

import "fmt"

//生产者
func producer(tasks chan int, taskSize int) {

	for i := 0; i < taskSize; i++ {
		tasks <- i
	}
}

//消费者
func consumer(tasks chan int, taskSize int, done chan bool) {

	for i := 0; i < taskSize; i++ {

		task := <-tasks
		fmt.Println(task)
	}
	done <- true
}

//主函数
func main() {
	fmt.Println("chan demo")

	tasks := make(chan int)
	done := make(chan bool)

	var taskSize = 10

	go producer(tasks, taskSize)
	go consumer(tasks, taskSize, done)

	<-done

	fmt.Println("done")
}
