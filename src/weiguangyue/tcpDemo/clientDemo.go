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

	min_len := 4 + 4 + 4
	array := make([]byte, 0)
	for {
		arr := make([]byte, 4)
		cnt, err := conn.Read(arr)
		if err != nil {

			log.Printf("read error:%s\n", err.Error())
			os.Exit(-1)
		}
		if cnt > 0 {

			log.Printf("read count:%d\n", cnt)

			array = append(array, arr...)
			if len(array) >= min_len {

				actual_length_arr := array[0:4]
				var actual_length uint32 = uint32(binary.LittleEndian.Uint32(actual_length_arr))

				split_length := actual_length + uint32(min_len)
				packet := array[0:split_length]
				process_packet(packet)

				array = make([]byte, 0)
			}

		} else if cnt == 0 {

			//log.Printf("read again\n")
		} else {

			log.Printf("read error\n")
			os.Exit(-1)
		}
	}
	conn.Close()
}

func process_packet(array []byte) {

	var seq uint32 = binary.LittleEndian.Uint32(array[0:4])
	var length uint32 = binary.LittleEndian.Uint32(array[4:8])
	var packet_type uint32 = binary.LittleEndian.Uint32(array[8:12])

	log.Printf("process packet.length:%d,seq:%d,packet_type:%d.\n", length, seq, packet_type)
}
