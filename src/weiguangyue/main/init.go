package main

import "log"

func init() {
	log.SetPrefix("")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Printf("init log complate!!!\n")
}
