package main

type RequestType uint32

const{
	RequestType_Ping RequestType = 0
	RequestType_Pong RequestType = 1
	RequestType_SendMsg RequestType = 2
	RequestType_RespMsg RequestType = 3
}

type RequestResponseHandler interface{

	sendPing()
	sendPong()
	sendMsg()
	sendRespMsg()

}
