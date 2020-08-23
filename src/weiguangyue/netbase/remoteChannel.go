package netbase

import (
	"encoding/binary"
	"log"
	"net"
)

type Channel interface {
	SendPing()
	SendPong(seq uint32)
	SendMsg(msg []byte)
	SendRespMsg(seq uint32)

	Id() uint32
	IsClose() bool
	Close()
}

type MessageHandler interface {
	Handle(channel Channel, msg []byte)
}

/**
type RemoteChannelListener interface {
	OnConnect()
	OnClose()
}
**/

type ServerChannel struct {
	cls            bool
	cid            uint32
	ln             net.Listener
	clients        map[uint32]*ClientChannel
	messageHandler MessageHandler
}

func (serverChannel *ServerChannel) setMessageHandler(messageHandler MessageHandler) {

	serverChannel.messageHandler = messageHandler
}

func (serverChannel *ServerChannel) nextCid() uint32 {

	serverChannel.cid++
	return serverChannel.cid
}

func (serverChannel *ServerChannel) IsClose() bool {

	return serverChannel.cls
}

func (serverChannel *ServerChannel) Close() {

	if serverChannel.cls {
		return
	}

	for cid, clientChannel := range serverChannel.clients {

		log.Printf("close client[id=%d]\n", cid)
		clientChannel.Close()
		delete(serverChannel.clients, cid)
	}

	serverChannel.ln.Close()

	serverChannel.cls = true
}

func doAccpet(serverChannel *ServerChannel) {

	for {
		log.Printf("wait accept new connection \n")
		conn, err := serverChannel.ln.Accept()
		if err != nil {
			// handle error
			log.Printf("accept error:%s\n", err.Error())
			conn.Close()
		} else {

			cid := serverChannel.nextCid()

			log.Printf("accept new client[id=%d]\n", cid)

			clientChannel := &ClientChannel{
				seq:            0,
				cls:            false,
				cid:            cid,
				conn:           conn,
				messageHandler: serverChannel.messageHandler,
			}

			serverChannel.clients[cid] = clientChannel

			clientChannel.readClientChannel()
		}
	}
}

func Bind(address string, messageHandler MessageHandler) (serverChannel *ServerChannel, err error) {

	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	sc := &ServerChannel{
		cid:            0,
		cls:            false,
		ln:             ln,
		clients:        make(map[uint32]*ClientChannel),
		messageHandler: messageHandler,
	}
	go doAccpet(sc)

	return sc, nil
}

type ClientChannel struct {
	seq            uint32
	cls            bool
	cid            uint32
	conn           net.Conn
	messageHandler MessageHandler
}

func (clientChannel *ClientChannel) readClientChannel() {

	go doReadClientChannelRoutine(clientChannel)
}

func doReadClientChannelRoutine(clientChannel *ClientChannel) {

	err := doReadClientChannel(clientChannel)
	if err.Error() == "EOF" {
		log.Printf("close")
	} else {
		log.Printf("read error: %s\n", err.Error())
	}

}

func doReadClientChannel(clientChannel *ClientChannel) error {

	conn := clientChannel.conn

read_continue_loop:
	for {
		var length uint32 = 0
		var need_to_read_size int = 4

		var buf []byte = make([]byte, 0)

		for {
			buf_to_read := make([]byte, need_to_read_size)

		read_length_continue:
			cnt, err := conn.Read(buf_to_read)
			if err != nil {
				return err
			}
			if cnt > 0 {
				need_to_read_size -= cnt

				if need_to_read_size > 0 {
					temp := buf_to_read[0:cnt]
					buf = append(buf, temp...)
					goto read_length_continue
				} else if need_to_read_size == 0 {

					temp := buf_to_read[0:cnt]
					buf = append(buf, temp...)

					length = binary.LittleEndian.Uint32(buf)
					goto read_data_loop
				}
			} else if cnt == 0 {
				goto read_length_continue
			} else {
				return nil
			}
		}
	read_data_loop:

		need_to_read_size = int(length)
		buf = make([]byte, 0)
		for {

		read_data_conitnue:
			buf_to_read := make([]byte, need_to_read_size)

			cnt, err := conn.Read(buf_to_read)
			if err != nil {
				return err
			} else {
				if cnt > 0 {

					need_to_read_size -= cnt

					if need_to_read_size > 0 {
						temp := buf_to_read[0:cnt]
						buf = append(buf, temp...)

						goto read_data_conitnue

					} else if need_to_read_size == 0 {
						temp := buf_to_read[0:cnt]
						buf = append(buf, temp...)

						clientChannel.messageHandler.Handle(clientChannel, buf)

						goto read_continue_loop
					}
				} else if cnt == 0 {
					//read again ???
				} else {
					return nil
				}
			}
		}
	}
	return nil
}

func (clientChannel ClientChannel) Id() uint32 {

	return clientChannel.cid
}

func (clientChannel *ClientChannel) nextSeq() uint32 {

	clientChannel.seq++
	return clientChannel.seq
}

func (clientChannel *ClientChannel) IsClose() bool {

	return clientChannel.cls
}

func (clientChannel *ClientChannel) Close() {

	if clientChannel.cls {

		return
	}

	clientChannel.conn.Close()
	clientChannel.cls = true
	log.Printf("close client[id=%d]\n", clientChannel.cid)
}

func (clientChannel *ClientChannel) SendPing() {

	SendPing(clientChannel.conn, clientChannel.nextSeq())
}

func (clientChannel *ClientChannel) SendPong(seq uint32) {

	SendPong(clientChannel.conn, seq)
}

func (clientChannel *ClientChannel) SendMsg(msg []byte) {

	SendMsg(clientChannel.conn, clientChannel.nextSeq(), msg)
}

func (clientChannel *ClientChannel) SendRespMsg(seq uint32) {

	SendRespMsg(clientChannel.conn, seq)

}
func Connect(remoteAddress string, messageHandler MessageHandler) (clientChannel *ClientChannel, err error) {

	log.Printf("connect %s\n", remoteAddress)
	conn, err := net.Dial("tcp", remoteAddress)
	if err != nil {
		log.Printf("connect error:%s\n", err.Error())
		return nil, err
	}
	log.Printf("connect %s success \n", remoteAddress)

	cc := &ClientChannel{
		seq:            0,
		cls:            false,
		cid:            0,
		conn:           conn,
		messageHandler: messageHandler,
	}

	cc.readClientChannel()

	return cc, nil
}
