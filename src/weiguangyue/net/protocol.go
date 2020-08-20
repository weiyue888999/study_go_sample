package main

import (
	"encoding/binary"
	"net"
)

type RequestType uint32

const (
	RequestType_Ping    RequestType = 0
	RequestType_Pong    RequestType = 1
	RequestType_SendMsg RequestType = 2
	RequestType_RespMsg RequestType = 3
)

func SendPing(conn net.Conn, seq uint32) {

	sendTo := make([]byte, 0)

	var length uint32 = 0
	//seq
	{
		length += 4
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, seq)
		sendTo = append(sendTo, arr...)
	}

	//type
	{
		length += 4
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(RequestType_Ping))
		sendTo = append(sendTo, arr...)
	}

	//length
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(length))
		sendTo = append(arr, sendTo...)
	}
	conn.Write(sendTo)
}

func SendPong(conn net.Conn, seq uint32) {
	sendTo := make([]byte, 0)

	var length uint32 = 0
	//seq
	{
		length += 4
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, seq)
		sendTo = append(sendTo, arr...)
	}

	//type
	{
		length += 4
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(RequestType_Pong))
		sendTo = append(sendTo, arr...)
	}

	//length
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(length))
		sendTo = append(arr, sendTo...)
	}
	conn.Write(sendTo)

}

func SendMsg(conn net.Conn, seq uint32, msg string) {
	sendTo := make([]byte, 0)

	var length uint32 = 0
	//seq
	{
		length += 4
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, seq)
		sendTo = append(sendTo, arr...)
	}

	//type
	{
		length += 4
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(RequestType_SendMsg))
		sendTo = append(sendTo, arr...)
	}

	{
		length += len(msg)
		arr := []byte(msg)
		sendTo = append(sendTo, arr...)
	}

	//length
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(length))
		sendTo = append(arr, sendTo...)
	}
	conn.Write(sendTo)

}

func SendRespMsg(conn net.Conn, seq uint32) {
	sendTo := make([]byte, 0)

	var length uint32 = 0
	//seq
	{
		length += 4
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, seq)
		sendTo = append(sendTo, arr...)
	}

	//type
	{
		length += 4
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(RequestType_RespMsg))
		sendTo = append(sendTo, arr...)
	}

	//length
	{
		arr := make([]byte, 4)
		binary.LittleEndian.PutUint32(arr, uint32(0))
		sendTo = append(arr, sendTo...)
	}
	conn.Write(sendTo)

}
