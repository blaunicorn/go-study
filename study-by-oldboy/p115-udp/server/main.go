// UDP server
package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 40000,
	})
	if err != nil {
		fmt.Println("listen UDP failed,err:", err)

	}
	defer conn.Close()
	// need not establish a connection,Send data directly
	var data [1024]byte
	for {
		n, addr, err := conn.ReadFromUDP(data[:]) // array becomes slice
		if err != nil {
			fmt.Println("read from UDP failed,err:", err)
			return
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		// fmt.Println(data[:n])
		// fmt.Println(string(data[:n]))
		// convert the received data to uppercase. slice to string,  then string to uppercase.
		reply := strings.ToUpper(string(data[:n]))
		// send data
		conn.WriteToUDP([]byte(reply), addr)
	}
}
