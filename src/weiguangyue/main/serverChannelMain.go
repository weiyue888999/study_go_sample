package main

import (
	netbase "weiguangyue/netbase"
)

func main() {

	serverChannel := netbase.Bind("localhost:8080")
	serverChannel.Close()
}
