package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int

	// 在线用户的列表
	OnLineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{Ip: ip, Port: port, OnLineMap: make(map[string]*User), Message: make(chan string)}
	return server
}

// 监听Message广播消息channel的goroutine，一旦有消息，就发送给全部的在线user
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		// 将msg发送给全部的在线User
		this.mapLock.Lock()
		for _, cli := range this.OnLineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// 广播消息的方法
func (this *Server) BroadCast(user *User, msg string) {

	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
	// fmt.Println(user, sendMsg)
}

func (this *Server) Handler(conn net.Conn) {
	// 当前链接的业务
	fmt.Println("链接建立成功...")

	//创建用户user
	user := NewUser(conn, this)
	// fmt.Println(user, conn)

	// 把用户上线的业务逻辑放到用户里去
	// // 用户上线，将用户加入onlineMap中
	// this.mapLock.Lock()
	// this.OnLineMap[user.Name] = user
	// this.mapLock.Unlock()

	// //广播当前用户上线消息
	// this.BroadCast(user, "已上线")

	// 替换成user方法
	user.Online()

	//监听用户是否活跃的channel
	isLive := make(chan bool)

	// //接受客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				// this.BroadCast(user, "下线")
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
				return
			}
			//提取用户的消息，去除“\n"
			msg := string(buf[:n-1])

			// 将得到的消息进行广播
			this.BroadCast(user, msg)
			// 用户针对msg进行处理
			user.DoMessage(msg)

			// 用户的任意消息，代表当前用户是一个活跃的
			isLive <- true
		}
	}()

	// 当前handler阻塞
	for {
		select {
		case <-isLive:
			// 当前用户是活跃的，应当重置定时器
		case <-time.After(time.Second * 300):
			// 已经超时，
			// 将当前的user强制的关闭

			user.SendMsg("你被踢了")

			// 销毁用户句柄
			close(user.C)
			// this.mapLock.Lock()
			// delete(this.OnLineMap, user.Name)
			// this.mapLock.Unlock()
			//关闭当前链接,因为关闭conn链接，服务器从buf读取数据长度为0，将触发offline清理map，就不需要手动写了。
			conn.Close()
			// 退出当前的Handler
			return
		}
	}

}

// 启动服务器的接口
func (this *Server) Start() {

	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
	}
	fmt.Println("server start")
	// close listen socket
	defer listener.Close()

	// 启动监听messager的goroutine
	go this.ListenMessager()

	//accept
	// do handler
	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		go this.Handler(conn)
	}

	// close listen socket

}

// go build -o server.exe main.go server.go user.go
