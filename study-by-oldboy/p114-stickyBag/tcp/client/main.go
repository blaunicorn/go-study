// tcp client
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	proto "p114-stickyBag/tcp/procotol"
	"strings"
)

func main() {
	fmt.Println("start the client...")

	ClientBase()

}

func ClientBase() {
	// 1.	Establish a connection with the server.
	conn, err := net.Dial("tcp", "192.168.101.32:8080")
	if err != nil {
		fmt.Println("dial 192.168.101.32:8080 failed,err:", err.Error())
		return
	}
	defer conn.Close() // close connection
	// 2. send data
	// conn.Write([]byte("hello wangye!")) // 简单数据

	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("please say:")
		// fmt.Scanln(&msg) // 遇到空格就会结束
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "EXIT" {
			return
		}
		// _, err = conn.Write([]byte(inputInfo)) // 发送数据
		// // 调用协议编码数据
		b, err := proto.Encode(inputInfo)
		if err != nil {
			fmt.Println("编码错误：", err)
			return
		}
		_, err = conn.Write(b) // 发送数据
		if err != nil {
			fmt.Println("写入错误", err)
			return
		}

		// 读取server端发送过来的数据
		// buf := [512]byte{}
		// n, err := conn.Read(buf[:])
		// if err != nil {
		// 	fmt.Println("rece failed,err:", err)
		// 	return
		// }
		// fmt.Println("rece server :", string(buf[:n]))
	}
}

// 	inputReader := bufio.NewReader(os.Stdin)
// 	fmt.Println("please input your name:")
// 	clientName, _ := inputReader.ReadString('\n')
// 	inputClientName := strings.Trim(clientName, "\n")
// 	//send info to server until Quit
// Loop:
// 	for {
// 		fmt.Println("what do ou send to the server? Type Q to quit。")
// 		content, _ := inputReader.ReadString('\n')
// 		inputContent := strings.Trim(content, "\n")
// 		println(inputContent == "Q")
// 		if inputContent == "Q" {
// 			println(inputContent)
// 			break Loop
// 		}

// 		_, err := conn.Write([]byte(inputClientName + " say " + inputContent))
// 		if err != nil {
// 			fmt.Println("error write:", err.Error())
// 			return
// 		}
// 	}
// }

// 心跳包示例
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"os"
// 	"time"
// )

// func main() {
// 	//开一个goroutine，做连接，并启动接收，如果连接断开，接收报错，则尝试恢复连接
// 	go Link()
// 	//来一个goroutine，做心跳，如果不发心跳，拔掉网线不会触发接收异常
// 	go BeatHeart()
// 	//在启动的goroutine中做发送操作
// 	input := bufio.NewScanner(os.Stdin)
// 	for {
// 		input.Scan()
// 		str := input.Text()
// 		fmt.Println("你输入的是：", str)
// 		sendBuf := []byte(str)
// 		if conn == nil {
// 			fmt.Println("conn is nil")
// 			continue
// 		}
// 		n, err := conn.Write(sendBuf[:])
// 		if err != nil {
// 			fmt.Println(err)
// 		} else {
// 			fmt.Println("send:", n)
// 		}
// 	}
// }

// //全局tcp连接对象
// var conn net.Conn

// //负责连接以及连接恢复
// func Link() {
// 	hostInfo := "192.168.101.32:8080"
// 	for {
// 		var err error
// 		conn, err = net.Dial("tcp", hostInfo)
// 		fmt.Print("connect (")
// 		if err != nil {
// 			fmt.Println(") fail", err)
// 		} else {
// 			fmt.Println(") ok")
// 			defer func() {
// 				conn.Close()
// 				conn = nil
// 			}()
// 			doTask(conn)
// 		}
// 		time.Sleep(3 * time.Second)
// 	}
// }

// //心跳 每8秒发送一个包
// func BeatHeart() {
// 	for {
// 		if conn != nil {
// 			conn.Write([]byte("beatHeart"))
// 		}
// 		time.Sleep(time.Second * 8)
// 	}
// }

// //接收
// func doTask(conn net.Conn) {
// 	var buf [128]byte
// 	for {
// 		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
// 		n, err := conn.Read(buf[:])
// 		if err != nil {
// 			fmt.Println("接收错误，进行重连：", err)
// 			break
// 		}
// 		fmt.Println("接收到：", string(buf[:n]))
// 	}
// }
