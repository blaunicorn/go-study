package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	Name       string
	conn       net.Conn
	flag       int // 当前客户端的模式
}

// 创建客户端对象
// 链接服务器
// 返回客户端
func NewClient(serverIp string, serverPort int) *Client {
	client := &Client{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial error", err)
		return nil
	}
	client.conn = conn
	return client
}

// 设置命令行工具
var serverIp string
var serverPort int

// 帮助说明
// ./client -h
// 指定端口和ip运行
// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器ip地址（默认是127.0.0.1）")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口（默认是8888）")
}

// 显示菜单
func (client *Client) menu() bool {
	var flag int //这个应该是客户端的属性
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更改用户名")
	fmt.Println("0.退出")
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("...请输入合法范围内的数字")
		return false
	}
}

func (client *Client) UpdateName() bool {
	fmt.Println(">>>请输入用户名：")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn.Write err:", err)
		return false
	}
	return true
}

// 处理server回应的消息，直接显示到标准输出即可,并行处理
func (client *Client) DealResponse() {
	io.Copy(os.Stdout, client.conn)
	// 一旦client conn 有数据就直接copy到stdout的标准输出上，
	// 等同于
	// for {
	// 	buf := make()
	// 	client.conn.Read(buf)
	// 	fmt.Println(buf)
	// }
}
func (client *Client) PublicChat() {
	// 提示用户输入消息
	var chatMsg string
	fmt.Println(">>>请输入聊天内容，exit退出")
	fmt.Scanln(&chatMsg)
	for chatMsg != "exit" {
		//发给服务器
		//如果消息为空则不发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write err:", err)
				break
			}

		}
		chatMsg = ""

		fmt.Println(">>>请输入聊天内容，exit退出")
		fmt.Scanln(&chatMsg)
	}

}

// 查询在线yonghu
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return
	}
}

func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string
	client.SelectUsers()
	fmt.Println("请输入聊天对象[用户名]，exit退出：")
	fmt.Scanln(&remoteName)
	for remoteName != "exit" {
		fmt.Println(">>>请输入消息内容，exit退出")
		fmt.Scanln(&chatMsg)
		for chatMsg != "exit" {
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err:", err)
					break
				}

			}
			chatMsg = ""
			fmt.Println(">>>请输入聊天内容，exit退出")
			fmt.Scanln(&chatMsg)
		}
		client.SelectUsers()
		fmt.Println("请输入聊天对象[用户名]，exit退出：")
		fmt.Scanln(&remoteName)
	}
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() != true {

		}
		// 根据不同的模式处理不同的业务
		switch client.flag {
		case 1:
			// 公聊模式
			fmt.Println("公聊模式选择")
			client.PublicChat()
			break
		case 2:
			// 私聊模式
			fmt.Println("私聊模式选择")
			break
		case 3:
			// 更改用户名
			fmt.Println("更新用户名选择")
			client.UpdateName()
			break
		}
	}
}

func main() {
	// 命令行解析
	flag.Parse()

	client := NewClient(serverIp, serverPort)
	// client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>>链接服务器失败...")
		return
	}
	// 单独开启一个goroutine 去处理server的回执消息
	go client.DealResponse()

	fmt.Println(">>>>链接服务器成功...")

	// 启动客户端的业务
	// select {}
	client.Run()
}
