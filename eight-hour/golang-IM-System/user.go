package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
	// 当前用户是属于那个server，好在用户权限内判断
	server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}

	// 启动监听当前user channel消息的go routine
	go user.ListenMessage()

	return user
}

// 用户的上线业务
func (this *User) Online() {
	// 用户上线，将用户加入onlineMap中
	this.server.mapLock.Lock()
	this.server.OnLineMap[this.Name] = this
	this.server.mapLock.Unlock()

	//广播当前用户上线消息
	this.server.BroadCast(this, "已上线")
}

// 用户的下线业务
func (this *User) Offline() {
	// 用户下线，将用户从onlineMap中删除
	this.server.mapLock.Lock()
	delete(this.server.OnLineMap, this.Name)
	this.server.mapLock.Unlock()

	//广播当前用户下线消息
	this.server.BroadCast(this, "已下线")
}

func (this *User) SendMsg(msg string) {
	// 给当前用户对应的客户端发送消息
	this.conn.Write([]byte(msg))
	// this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前用户都有哪些
		this.server.mapLock.Lock()
		for _, user := range this.server.OnLineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 消息格式：rename|张三
		newName := msg[8:]
		// newName:=strings.Split(msg,"|")[1]
		//判断name是否存在
		_, ok := this.server.OnLineMap[newName]
		if ok {
			this.SendMsg("当前用户名被使用")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnLineMap, this.Name)
			this.server.OnLineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			// windows下可能是 "\r\n"
			this.SendMsg("您已经更新用户名：" + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 消息格式： to|张三|消息

		// 1. 获对方的用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式不正确，请使用\"to|张三|消息\"格式。\n")
		}

		// 2. 根据用户名，得到对方user对象
		remoteUser, ok := this.server.OnLineMap[remoteName]
		if !ok {
			this.SendMsg("该用户名不存在")
			return
		}
		// 3. 获取消息内容，通过对方的user对象将消息内容发送过去。
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容，请重发\n")
			return
		}
		remoteUser.SendMsg(this.Name + "对您说：" + content)
		return
	} else {
		this.server.BroadCast(this, msg)
	}

}

// 监听当前User channel的方法，一旦有消息，就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {

		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
		fmt.Println(this.Name + ":user listennmessage")
	}
}
