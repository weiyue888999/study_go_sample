package main

import "fmt"
import "strconv"

//生产event
func producerEvent(events chan string, done chan int) {

	for i := 0; i < 10; i++ {
		str := strconv.Itoa(i)
		events <- str
	}

	done <- 1
}

//生产计算任务
func producerTask(tasks chan int, done chan int) {

	for i := 0; i < 10; i++ {
		tasks <- i
	}

	done <- 1
}

func consume(tasks chan int, events chan string) {

	fmt.Println("start consumer")

	for {
		select {
		case i := <-tasks:
			fmt.Println(i)
		case j := <-events:
			fmt.Println(j)
		}
	}
}

func main() {

	var tasks chan int = make(chan int)
	var events chan string = make(chan string)
	var done chan int = make(chan int)

	go producerTask(tasks, done)
	go producerEvent(events, done)
	go consume(tasks, events)

	for i := 0; i < 2; i++ {
		<-done
	}

	defer close(tasks)
	defer close(events)
	defer close(done)
}
