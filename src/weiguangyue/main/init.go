package main

import (
	"log"
	"fmt"
)

func init() {

	fmt.Println("init...")
	log.SetPrefix("")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Printf("init log complate!!!\n")
}
