// tcp server
package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("start the server...")
	// 1.Local port start service
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println(" start tcp server on 0.0.0.0:8080 error, err:", err.Error())
		return // end function
	}
	defer listener.Close()
	// 2.Wait for the client  to request to establish a connection
	for {
		conn, err := listener.Accept() // Establish connection
		if err != nil {
			fmt.Println("error accepting", err.Error())
			continue
		}
		// 3. Communicating with clients
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Println("连接断开...")
	}()
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf[:])
		if err == io.EOF {
			fmt.Println(" reading from conn io.EOF,err:", err.Error())
			break
		}
		if err != nil {
			fmt.Println(" reading from conn failed,err:", err.Error())
			break
		}
		fmt.Printf("received data: %v/n", string(buf[:len]))
		fmt.Println()
		fmt.Printf("received data: %v/n", buf[:len])
		fmt.Println()
		// replay demo
		var msg string
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("please replay:")
			// fmt.Scanln(&msg) // 遇到空格就会结束
			text, _ := reader.ReadString('\n')
			msg = strings.TrimSpace(text)
			if msg == "exit/n" {
				break
			}
			conn.Write([]byte(msg))
		}
	}
}
