package main

import (
	"fmt"
	"net"
)

func main() {
	udpAddr := net.UDPAddr{
		Port: 6969,
		IP:   net.ParseIP("127.0.0.1"),
	}
	listener, err := net.ListenUDP("udp", &udpAddr)
	if err != nil {
		panic(err)
	}

	packet := make([]byte, 64)
	fmt.Println("UDP Listening on 6969")
	for {
		n, remoteAddress, err := listener.ReadFromUDP(packet)
		if err != nil {
			panic(err)
		}
		fmt.Printf(
			"Message from: %v, of length: %v, Containing: %s\n",
			remoteAddress,
			n,
			string(packet),
		)
		_, err = listener.WriteToUDP([]byte("Hey ;)"), remoteAddress)
		if err != nil {
			panic(err)
		}
	}
}
