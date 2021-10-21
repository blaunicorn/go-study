// UDP client
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// client dircet dialing
func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("failed to connect to the server.err:", err)
		return
	}
	defer socket.Close()
	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("please input thecontent:")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg)) //
		// receive data returned by the server
		n, remoteAddr, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("read reply msg failed,err:", err)
			return
		}
		// fmt.Println("received data returned:", string(reply[:n]))
		fmt.Printf("recv:%v addr:%v count:%v\n", string(reply[:n]), remoteAddr, n)

	}
}
