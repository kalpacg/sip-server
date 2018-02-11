package main

import (
	"github.com/uname/siputility"
)

func main() {
	network := "udp"
	address := "192.168.1.34:5060"

	udpConn := siputility.StartConnection(network, address)

	buf := make([]byte, 1024)

	for {
		siputility.ReadData(udpConn, buf)
	}
}
