package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModeAppend)

	if err != nil {

		fmt.Println(err)
		os.Exit(-1)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	buf := make([]byte, 100)

	_, err1 := reader.Read(buf)
	if err1 != nil {

		fmt.Println(err1)
		os.Exit(-1)
	}

	var str string = string(buf)

	fmt.Println(str)
}
