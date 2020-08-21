package netbase

import "net"

type ConnectionHandler interface {
	Handle(conn net.Conn, seq uint32, packet_type uint32, data []byte)
}

type PingHandler struct {
}

func (h *PingHandler) Handle(conn net.Conn) {

}

type PongHandler struct {
}

func (h *PongHandler) Handle(conn net.Conn) {

}
