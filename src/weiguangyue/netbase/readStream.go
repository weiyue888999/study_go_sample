package netbase

import (
	"encoding/binary"
	"log"
	"net"
	"os"
)

func ReadConn( //
	conn net.Conn, //
	handlerMapping map[uint32]func(seq uint32, packet_type RequestType, data []byte, conn net.Conn)) {
	for {

		//read length
		read_total_cnt := 0
		length_array := make([]byte, 4)
		var length uint32 = 0
	read_head:
		cnt, err := conn.Read(length_array)
		if err != nil {

			log.Printf("read error:%s\n", err.Error())
			os.Exit(-1)
		}
		if cnt > 0 {

			log.Printf("read cnt:%d\n", cnt)
			read_total_cnt += cnt

			if read_total_cnt != 4 {
				goto read_head
			} else {
				length = binary.LittleEndian.Uint32(length_array)
				log.Printf("packet length:%d\n", length)
				goto read_body
			}
		} else if cnt == 0 {

			//log.Printf("read again\n")
		} else {

			log.Printf("read error\n")
			os.Exit(-1)
		}
	read_body:
		read_total_cnt = 0
		packet_array := make([]byte, length)
	read_body_continue:
		cnt, err = conn.Read(packet_array)
		if err != nil {
			log.Printf("read error:%s\n", err.Error())
			os.Exit(-1)
		} else {
			if cnt > 0 {

				log.Printf("read cnt:%d\n", cnt)

				read_total_cnt += cnt
				if uint32(read_total_cnt) == length {
					process_packet(length, packet_array, conn, handlerMapping)
				} else {
					goto read_body_continue
				}
			} else if cnt == 0 {

				//read again ???
			} else {
				log.Printf("read error\n")
				os.Exit(-1)
			}
		}
	}
}

func process_packet( //
	length uint32, //
	array []byte, //
	conn net.Conn, //
	handlerMapping map[uint32]func(seq uint32, packet_type RequestType, data []byte, conn net.Conn)) { //

	var seq uint32 = binary.LittleEndian.Uint32(array[0:4])
	var packet_type uint32 = binary.LittleEndian.Uint32(array[4:8])
	var data []byte = make([]byte, 0)
	if length > 4+4 {
		data = array[8:length]
	}
	fun, found := handlerMapping[uint32(packet_type)]

	if found {
		log.Printf("process packet,seq:%d,packet_type:%d\n", seq, packet_type)

		fun(seq, RequestType(packet_type), data, conn)
	} else {
		log.Printf("unknow packet type:%d\n", packet_type)
	}
}
