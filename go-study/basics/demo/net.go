package demo

import (
	"fmt"
	"net"
)

func NetClass() {
	fmt.Println("net tcp:")

	/*
		tcp
		客户端 net.DialTcp("tcp",nill,tcpAddr)
		服务端 net.ListenTcp("协议"，addr)
		       net.ResolveTCPAddr("协议","端口")

		http
		客户端 get post ...
		服务端

		rpc
		客户端
		服务端
	*/
	//创建服务端
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "8888") // 创建地址
	listener, _ := net.ListenTCP("tcp", tcpAddr)
	// 循环监听，调用,以为不仅仅有一个连接，会有多个连接
	for {
		TCPConn, err := listener.AcceptTCP() //创建tcp服务端
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(TCPConn)
	}
}

func handleConnection(conn *net.TCPConn) {
	// buf := make([]byte, 1024)
	// n, _ := conn.Read(buf)
	// fmt.Println(conn.RemoteAddr().String()+" Connected ", n)

	// 让服务器一直读信息
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(conn.RemoteAddr().String()+" Connected ", n)

		// 服务器 给 客户端 返回数据
		str := "收到了:" + string(buf[0:n])
		conn.Write([]byte(str))
	}
}
