package demo

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func NetClientClass() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "8888") // 创建地址
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("net-client:", err)
		return
	}
	conn.Write([]byte("it is me"))
	// 命令行操作
	reader := bufio.NewReader(os.Stdin)
	// 用for循环 保持长连接
	for {
		bytes, _, _ := reader.ReadLine()
		conn.Write(bytes)

		// 读取 服务器 发回的数据
		readBuf := make([]byte, 1023)
		n, _ := conn.Read(readBuf)
		fmt.Println("服务器返回数据为：", string(readBuf[0:n]))
	}

}
