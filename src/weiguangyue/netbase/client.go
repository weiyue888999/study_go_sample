package netbase

import (
	"log"
	"net"
	"strconv"
)

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
