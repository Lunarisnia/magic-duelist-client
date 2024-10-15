package magicp

import (
	"encoding/json"
	"net"
)

func Listen(udpAddr net.UDPAddr, callback func(snapshot *SnapshotProtocol)) error {
	listener, err := net.ListenUDP("udp", &udpAddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	receivedPacket := make([]byte, 2048)
	serverSnapshot := SnapshotProtocol{}
	for {
		n, _, err := listener.ReadFromUDP(receivedPacket)
		if err != nil {
			return err
		}

		err = json.Unmarshal(receivedPacket[:n], &serverSnapshot)
		if err != nil {
			return err
		}

		callback(&serverSnapshot)
	}
}
