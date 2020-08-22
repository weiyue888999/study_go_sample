package main

import (
	netbase "weiguangyue/netbase"
)

func main() {

	clientChannel := netbase.Connect("localhost:8080")
	clientChannel.Write("hello,world!!!\n")
	clientChannel.Close()
}
