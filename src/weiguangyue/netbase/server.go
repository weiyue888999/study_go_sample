package netbase

import (
	"log"
	"net"
	"time"
)

type TcpServer struct {
	nextId         int `json:server.port`
	clientCount    int
	ServerConfig   ServerConfig
	listener       net.Listener
	clients        map[int]*Client
	handlerMapping map[uint32]func(seq uint32, packet_type RequestType, data []byte, conn net.Conn)
}

func (tcpServer *TcpServer) getNextId() int {
	tcpServer.nextId++
	return tcpServer.nextId
}

func pingHandler(seq uint32, packet_type RequestType, data []byte, conn net.Conn) {

	log.Printf("receive ping\n")
}

func (tcpServer *TcpServer) processNewConnection(client *Client) {

	tcpServer.clientCount++
	tcpServer.clients[client.id] = client

	defer func(client *Client) {
		tcpServer.clientCount--
		delete(tcpServer.clients, client.id)
		log.Printf("leave client id=%d\n", client.id)
		client.Close()
	}(client)

	log.Printf("handle new client id=%d\n", client.id)

	ReadConn(client.conn, tcpServer.handlerMapping)
}

func listene_new_conn(ln net.Listener, tcpServer *TcpServer) {

	for {
		log.Printf("waiting new connection \n")
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			log.Printf("accept error:%s\n", err.Error())
			conn.Close()
		} else {
			client := &Client{
				id:   tcpServer.getNextId(),
				name: "client-",
				conn: conn,
			}
			if tcpServer.clientCount < tcpServer.ServerConfig.MaxConnectClient {

				log.Printf("accept new connection,clientId=%d\n", client.id)
				tcpServer.processNewConnection(client)
			} else {
				log.Printf("force close client cuase clientCount[%d] >= maxConnectClient[%d]\n",
					tcpServer.clientCount,
					tcpServer.ServerConfig.MaxConnectClient)
				client.Close()
			}
		}
	}
}

//先使用睡眠方式进行模拟定时器
func ping_timer(tcpServer *TcpServer) {

	for {
		time.Sleep(time.Duration(2) * time.Second)

		if len(tcpServer.clients) <= 0 {

			log.Printf("no clients for ping\n")
			continue
		}
		for id, client := range tcpServer.clients {

			log.Printf("ping client[id=%d]\n", id)
			SendPing(client.conn, client.NextSeq())
		}
	}
}

func (tcpServer *TcpServer) Startup() error {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Printf("listen err" + err.Error())
		return err
	}

	tcpServer.listener = ln
	tcpServer.clients = make(map[int]*Client)
	tcpServer.handlerMapping = make(map[uint32]func(seq uint32, packet_type RequestType, data []byte, conn net.Conn))
	tcpServer.handlerMapping[uint32(RequestType_Ping)] = pingHandler

	go listene_new_conn(ln, tcpServer)

	go ping_timer(tcpServer)

	return nil
}

func (tcpServer *TcpServer) Shutdown() {
	tcpServer.listener.Close()
}
