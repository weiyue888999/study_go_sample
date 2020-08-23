package netbase

import (
	"encoding/binary"
	"fmt"
	"net"
)

type RequestType uint32

const (
	RequestType_Ping    RequestType = 0
	RequestType_Pong    RequestType = 1
	RequestType_SendMsg RequestType = 2
	RequestType_RespMsg RequestType = 3
)

func Parse(packet []byte) (requestType RequestType, seq uint32, data []byte) {

	rtype_array := packet[0:4]
	rtype_value := binary.LittleEndian.Uint32(rtype_array)

	seq_array := packet[4:8]
	seq_value := binary.LittleEndian.Uint32(seq_array)

	if packet_len := len(packet); packet_len > 8 {

		data := packet[8:packet_len]
		return RequestType(rtype_value), seq_value, data

	} else {
		return RequestType(rtype_value), seq_value, nil
	}
}

func SendPing(conn net.Conn, seq uint32) {

	sendTo := make([]byte, 0)

	var length uint32 = 4 + 4 + 0

	//length
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(length))
		sendTo = append(sendTo, arr...)
	}

	//type
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(RequestType_Ping))
		sendTo = append(sendTo, arr...)
	}

	//seq
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(seq))
		sendTo = append(sendTo, arr...)
	}

	//msg
	{
		//no msg
	}

	cnt, err := conn.Write(sendTo)
	if err != nil {

		fmt.Printf("write error:%s\n", err.Error())
	} else {

		fmt.Printf("write cnt:%d\n", cnt)
	}
}

func SendPong(conn net.Conn, seq uint32) {

	sendTo := make([]byte, 0)

	var length uint32 = 4 + 4 + 0

	//length
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(length))
		sendTo = append(sendTo, arr...)
	}

	//type
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(RequestType_Pong))
		sendTo = append(sendTo, arr...)
	}

	//seq
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(seq))
		sendTo = append(sendTo, arr...)
	}

	//msg
	{
		//no msg
	}

	cnt, err := conn.Write(sendTo)
	if err != nil {

		fmt.Printf("write error:%s\n", err.Error())
	} else {

		fmt.Printf("write cnt:%d\n", cnt)
	}
}

func SendMsg(conn net.Conn, seq uint32, msg []byte) {

	sendTo := make([]byte, 0)
	msgLength := uint32(len(msg))

	var length uint32 = 4 + 4 + msgLength

	//length
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(length))
		sendTo = append(sendTo, arr...)
	}

	//type
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(RequestType_SendMsg))
		sendTo = append(sendTo, arr...)
	}

	//seq
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(seq))
		sendTo = append(sendTo, arr...)
	}

	//msg
	{
		sendTo = append(sendTo, msg...)
	}

	conn.Write(sendTo)

}

func SendRespMsg(conn net.Conn, seq uint32) {

	sendTo := make([]byte, 0)

	var length uint32 = 4 + 4 + 0

	//length
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(length))
		sendTo = append(sendTo, arr...)
	}

	//type
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(RequestType_RespMsg))
		sendTo = append(sendTo, arr...)
	}

	//seq
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(seq))
		sendTo = append(sendTo, arr...)
	}

	//msg
	{
		//no msg
	}

	conn.Write(sendTo)
}
