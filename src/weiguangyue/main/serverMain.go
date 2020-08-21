package main

import (
	"log"
	"os"
	"time"
	netbase "weiguangyue/netbase"
)

func main() {

	serverConfig := netbase.ServerConfig{
		MaxConnectClient: 100,
		Port:             8080,
	}

	tcpServer := &netbase.TcpServer{
		ServerConfig: serverConfig,
	}

	err := tcpServer.Startup()
	if err != nil {
		log.Printf("tcpServer startup err:%s\n", err.Error())
		os.Exit(-1)
	}

	for i := 0; i < 100; i++ {

		time.Sleep(time.Second * 1)
	}

	tcpServer.Shutdown()
}
