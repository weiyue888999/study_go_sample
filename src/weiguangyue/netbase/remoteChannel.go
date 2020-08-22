package netbase

import (
	"encoding/binary"
	"log"
	"net"
)

type Channel interface {
	Id() uint32
	Write(msg []byte)
	IsClose() bool
	Close()
}

type MessageHandler interface {
	handle(channel Channel, msg []byte)
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
				cid:  cid,
				conn: conn,
			}

			serverChannel.clients[cid] = clientChannel
		}
	}
}

func Bind(address string) (serverChannel *ServerChannel, err error) {

	ln, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	sc := &ServerChannel{
		cls: false,
		ln:  ln,
	}
	go doAccpet(sc)

	return sc, nil
}

type ClientChannel struct {
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
	log.Printf("read error: %s\n", err.Error())
}

func doReadClientChannel(clientChannel *ClientChannel) error {

	conn := clientChannel.conn

	for {
		var length uint32 = 0

		need_to_read_size := 4
		buf := make([]byte, 0)
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
					//FIXME one times to copy , not split
					temp := buf_to_read[0:cnt]
					buf = append(buf, temp...)
					goto read_length_continue
				}
				length = binary.LittleEndian.Uint32(buf)
				break
			} else if cnt == 0 {
				//log.Printf("read again\n")
			} else {
				log.Printf("read error\n")
			}
		}

		need_to_read_size = int(length)
		buf = make([]byte, 0)
		for {

			buf_to_read := make([]byte, need_to_read_size)

		read_data_continue:

			cnt, err := conn.Read(buf_to_read)
			if err != nil {
				return err
			} else {
				if cnt > 0 {

					need_to_read_size -= cnt

					if uint32(need_to_read_size) > 0 {
						temp := buf_to_read[0:cnt]
						buf = append(buf, temp...)
						goto read_data_continue
					}
					clientChannel.messageHandler.handle(clientChannel, buf)
					break

				} else if cnt == 0 {
					//read again ???
				} else {
					log.Printf("read error\n")
				}
			}
		}
	}
	return nil
}

func (clientChannel ClientChannel) Id() uint32 {

	return clientChannel.cid
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

func (clientChannel *ClientChannel) Write(msg []byte) {

	cnt, err := clientChannel.conn.Write(msg)
	if err != nil {
		log.Printf("wirte error:%s\n", err.Error())
	} else {
		log.Printf("wirte msg:%d\n", cnt)
	}
}

func Connect(remoteAddress string) (clientChannel *ClientChannel, err error) {

	log.Printf("connect %s\n", remoteAddress)
	conn, err := net.Dial("tcp", remoteAddress)
	if err != nil {
		log.Printf("connect error:%s\n", err.Error())
		return nil, err
	}
	log.Printf("connect %s success \n", remoteAddress)

	cc := &ClientChannel{
		cid:  0,
		conn: conn,
	}
	return cc, nil
}
