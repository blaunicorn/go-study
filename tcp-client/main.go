// 实现tcp客户端
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("start the client...")
	ClientBase()

}

func ClientBase() {
	// open connection:
	conn, err := net.Dial("tcp", "192.168.101.32:8080")
	if err != nil {
		fmt.Println("error dial:", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name:")
	clientName, _ := inputReader.ReadString('\n')
	inputClientName := strings.Trim(clientName, "\n")
	//send info to server until Quit
Loop:
	for {
		fmt.Println("what do ou send to the server? Type Q to quit。")
		content, _ := inputReader.ReadString('\n')
		inputContent := strings.Trim(content, "\n")
		println(inputContent == "Q")
		if inputContent == "Q" {
			println(inputContent)
			break Loop
		}

		_, err := conn.Write([]byte(inputClientName + " say " + inputContent))
		if err != nil {
			fmt.Println("error write:", err.Error())
			return
		}
	}
}
