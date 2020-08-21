package main

import (
	"log"
	"net"
	"os"
	"weiguangyue/netbase"
)

func pingHandler(seq uint32, packet_type netbase.RequestType, data []byte, conn net.Conn) {

	log.Printf("receive ping\n")
}

func main() {

	handlerMapping := make(map[uint32]func(seq uint32, packet_type netbase.RequestType, data []byte, conn net.Conn))
	handlerMapping[uint32(netbase.RequestType_Ping)] = pingHandler

	address := "localhost:8080"
	log.Printf("connect...%s\n", address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("connect error:%s\n", err.Error())
		os.Exit(-1)
	}
	log.Printf("connect success!!!\n")

	netbase.ReadConn(conn, handlerMapping)
}
