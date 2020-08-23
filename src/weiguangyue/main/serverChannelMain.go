package main

import (
	"fmt"
	"os"
	"time"
	netbase "weiguangyue/netbase"
)

type ServerChannelMessageHandler struct {
}

func (serverChannelMessageHandler ServerChannelMessageHandler) Handle(channel netbase.Channel, msg []byte) {

	rtype, seq, data := netbase.Parse(msg)

	if rtype == netbase.RequestType_Ping {

		channel.SendPong(seq)

	} else if rtype == netbase.RequestType_Pong {
		fmt.Printf("recive pong seq:%d\n", seq)

	} else if rtype == netbase.RequestType_SendMsg {

		fmt.Printf("recive rtype:%d,seq:%d, msg:%s\n", rtype, seq, string(data))
		channel.SendRespMsg(seq)

	} else if rtype == netbase.RequestType_RespMsg {

		fmt.Printf("recive respMsg seq:%d\n", seq)

	} else {

		fmt.Printf("unknow rtype:%d\n", rtype)
	}

}

func main() {

	serverChannelMessageHandler := ServerChannelMessageHandler{}
	serverChannel, err := netbase.Bind("localhost:8080", serverChannelMessageHandler)
	if err != nil {
		fmt.Printf("bind error:%s\n", err.Error())
		os.Exit(-1)
	}

	for i := 0; i < 100; i++ {

		time.Sleep(time.Second * 1)
	}

	serverChannel.Close()
}
