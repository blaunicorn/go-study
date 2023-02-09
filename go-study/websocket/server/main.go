package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
websoket.Upgrader {
	HandshakeTimeOut:0, 握手时间0为不限制
	ReadBufferSize:1024, 以字节为单位的IO缓冲区，如果缓冲区大小为零，则使用HTTP服务器分配的缓冲区
	WriteBufferSize: 1024, 以字节为单位的IO缓冲区，如果缓冲区大小为零，则使用HTTP服务器分配的缓冲区
	WriteBufferPool: nil, WriteBufferPool是用于写操作的缓冲池
	Subprotocols:nil, 按顺序指定服务器支持的协议
	Error:nil, 指定用于生成HTTP错误相应到达函数
	CheckOrigin:nil, 对过来的请求做校验用的
	EnableCompression: false, 指定服务器是否应尝试根据进行协议消息压缩
}

*/

// 创建websocket结构
var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// 为了客户端之间相互通信，需要存储每个conn链接，然后进行广播或单独通信
var conns []*websocket.Conn

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("11")
	conn, err := UP.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// 在每个客户端进入时 append
	conns = append(conns, conn)
	for {
		messageType, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		conn.WriteMessage(websocket.TextMessage, []byte("你说的是："+string(p)+"吧"))
		fmt.Println(messageType, string(p), e)
		// 广播告之
		for i := range conns {
			conns[i].WriteMessage(websocket.TextMessage, []byte("你说的是："+string(p)+"吧"))
		}
	}
	defer conn.Close()
	log.Println("服务关闭")
}

func main() {
	fmt.Println("server start")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)
}
