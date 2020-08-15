package main

import "fmt"
import "log"
import "os"

func main() {

	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_RDWR, os.ModeAppend)
	if err != nil {
		fmt.Printf("error open log file")
		fmt.Printf(err.Error())
		return
	}

	log := log.New(file, "log-demo", log.Llongfile)

	for i := 0; i < 10; i++ {
		log.Println("hello,go log!!!")
	}
}
