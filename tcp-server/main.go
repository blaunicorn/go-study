// 实现tcp服务器
package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("start the server...")
	// creater listener
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("error listening", err.Error())
		return // end function
	}
	// listening client request
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting", err.Error())
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("erro reading", err.Error())
			return
		}
		fmt.Printf("received data: %v/n", string(buf[:len]))
		fmt.Println()
		fmt.Printf("received data: %v/n", buf[:len])
		fmt.Println()
	}
}
