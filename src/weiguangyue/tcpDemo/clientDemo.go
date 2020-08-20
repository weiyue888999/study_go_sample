package main

import (
	"encoding/binary"
	"log"
	"net"
	"os"
	//	mynet "weiguangyue/net"
)

func main() {
	log.SetPrefix("")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	log.Printf("connect...\n")
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		// handle error
		log.Printf("connect error:%s\n", err.Error())
		os.Exit(-1)
	}
	log.Printf("connect success!!!\n")

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
					process_packet(length, packet_array)
				} else {
					goto read_body_continue
				}
			} else if cnt == 0 {

				//read again
			} else {
				log.Printf("read error\n")
				os.Exit(-1)
			}
		}
	}
	conn.Close()
}

func process_packet(length uint32, array []byte) {

	var seq uint32 = binary.LittleEndian.Uint32(array[0:4])
	var packet_type uint32 = binary.LittleEndian.Uint32(array[4:8])

	log.Printf("process packet.length:%d,seq:%d,packet_type:%d.\n", length, seq, packet_type)
}
