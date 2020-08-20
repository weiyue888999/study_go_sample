package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("connect...")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
		fmt.Println("connect error")
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	for i := 0; i < 10; i++ {
		b := []byte("weiguangyue")
		conn.Write(b)
	}

	conn.Close()
}
