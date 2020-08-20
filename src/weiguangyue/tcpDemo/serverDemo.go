package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"
	mynet "weiguangyue/net"
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
	nextId       int `json:server.port`
	clientCount  int
	serverConfig ServerConfig
	listener     net.Listener
	clients      map[int]*Client
}

func (tcpServer *TcpServer) getNextId() int {
	tcpServer.nextId++
	return tcpServer.nextId
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
			mynet.SendPing(client.conn, client.NextSeq())
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

	go listene_new_conn(ln, tcpServer)

	go ping_timer(tcpServer)

	return nil
}

func (tcpServer *TcpServer) Shutdown() {
	tcpServer.listener.Close()
}

type Client struct {
	id   int
	name string
	conn net.Conn
	seq  uint32
}

func (client *Client) NextSeq() uint32 {

	client.seq++
	return client.seq
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
