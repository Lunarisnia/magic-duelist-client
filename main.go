package main

import (
	"bufio"
	"fmt"
	"net"
)

// TODO: Do the start match API on the server, where you can assign each player to an ID
// TODO: Have their movement translated properly and is visibly moving well enough on both player

func main() {
	packet := make([]byte, 64)
	conn, err := net.Dial("udp", "127.0.0.1:6969")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	_, err = conn.Write([]byte("Hello, There!"))
	if err != nil {
		panic(err)
	}
	_, err = bufio.NewReader(conn).Read(packet)
	if err != nil {
		panic(err)
	}
	fmt.Println("Response: ", string(packet))
}
