package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"
	"weiguangyue/net/base"
)

type ServerConfig struct {
	maxConnectClient int
	port             int
}

func (serverConfig ServerConfig) String() string {

	str := "ServerConfig{maxConnectClient:" + strconv.Itoa(serverConfig.maxConnectClient) + "}"
	return str
}

type TcpServer struct {
	nextId       int
	clientCount  int
	serverConfig ServerConfig
	listener     net.Listener
}

func (tcpServer *TcpServer) getNextId() int {
	tcpServer.nextId++
	return tcpServer.nextId
}

func (tcpServer *TcpServer) processNewConnection(client *Client) {

	tcpServer.clientCount++

	defer func(client *Client) {
		tcpServer.clientCount--
		log.Printf("leave client id=%d\n", client.id)
		client.Close()
	}(client)

	log.Printf("handle new client id=%d\n", client.id)

	buf := make([]byte, 32)
	for {
		n, err := client.conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				log.Println("close by remote client id=%d\n", client.id)
			} else {
				log.Printf("read error:%s\n", err.Error())
			}
			return
		} else {
			if n > 0 {
				msg := string(buf[:n])
				log.Printf("read msg:%s\n", msg)
			} else if n == 0 {
				log.Printf("read no data\n")
			} else {
				//read eof
				return
			}
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

	go func(ln net.Listener) {

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
				if tcpServer.clientCount < tcpServer.serverConfig.maxConnectClient {

					log.Printf("accept new connection,clientId=%d\n", client.id)
					tcpServer.processNewConnection(client)
				} else {
					log.Printf("force close client cuase clientCount[%d] >= maxConnectClient[%d]\n",
						tcpServer.clientCount,
						tcpServer.serverConfig.maxConnectClient)
					client.Close()
				}
			}
		}

	}(ln)
	return nil
}

func (tcpServer *TcpServer) Shutdown() {
	tcpServer.listener.Close()
}

type Client struct {
	id   int
	name string
	conn net.Conn
}

func (client Client) String() string {

	str := "client{id=" + strconv.Itoa(client.id) + ",name=" + client.name + "}"
	return str
}

func (client *Client) Close() {
	client.conn.Close()
	log.Printf("close client{id=%d }\n", client.id)
}

func init() {
	log.SetPrefix("")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {

	base.Test()

	log.Println("init main")
	serverConfig := ServerConfig{
		maxConnectClient: 100,
		port:             8080,
	}

	tcpServer := &TcpServer{
		serverConfig: serverConfig,
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
