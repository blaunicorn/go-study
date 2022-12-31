package main

import (
	"fmt"
	"net"
)

func tcpDemo() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", ":8888")
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("listener:", err)
			return
		}
		go handleConnection(conn)
	}

}

func handleConnection(conn *net.TCPConn) {
	fmt.Println(conn.RemoteAddr().String() + "Connected")
}
