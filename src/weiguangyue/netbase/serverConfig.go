package netbase

import "strconv"

type ServerConfig struct {
	MaxConnectClient int
	Port             int
}

func (serverConfig ServerConfig) String() string {

	str := "ServerConfig{maxConnectClient:" + strconv.Itoa(serverConfig.MaxConnectClient) + "}"
	return str
}
