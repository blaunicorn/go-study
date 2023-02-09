package main

/*
	   1.1 软件架构
	   （ 1 ） 单一应用架构
	          当网站流量很小时，应用规模小时，只需一个应用，将所有功能都部署在一起，以减少部署服务器数量和成本。此时，用于简化增删改查工作量的数据访问框架 (ORM) 是关键。数 据库的处理时间影响应用的性能。
	   这种结构的应用适合小型系统，小型网站，或者企业的内部系统，用户较少，请求量不大，对请求的处理时间没有太高的要求。 将所有功能都部署到一个服务器，简单易用。开 发项目的难度低。
	   缺点：
	   1 、性能扩展比较困难
	   2 、不利于多人同时开发
	   3 、不利于升级维护
	   4 、整个系统的空间占用比较大
	   （ 2 ） 分布式服务架构
	           当应用越来越多，应用之间交互不可避免，将核心业务抽取出来，作为独立的服务，逐渐形成稳定的服务中心，使前端应用能更快速的响应多变的市场需求。此时，用于提高业务复用 及整合的 分布式服务框架 (RPC) 是关键。分布式系统将服务作为独立的应用，实现服务共享 和重用。
	   什么是分布式系统
	           分布式系统是若干独立计算机（服务器）的集合，这些计算机对于用户来说就像单个相关系统，分布式系统（ distributed system ）是建立在网络之上的服务器端一种结构。
	          分布式系统中的计算机可以使用不同的操作系统，可以运行不同应用程序提供服务，将服务分散部署到多个计算机服务器上。
	    RPC 【 Remote Procedure Call 】是指远程过程调用，是一种进程间通信方式，是一种技术思想，而不是规范。它允许程序调用另一个地址空间（网络的另一台机器上）的过程或函 数，而不用开发人员显式编码这个调用的细节。调用本地方法和调用远程方法一样。
	          RPC 的实现方式可以不同。例如 java 的 rmi, spring 远程调用等。
	          RPC 概念是在上世纪 80 年代由 Brue Jay Nelson( 布鲁 · 杰伊 · 纳尔逊 ) 提出。使用 PRC 可以将本地的调用扩展到远程调用（分布式系统的其他服务器）。
	   RPC 的特点
	   1. 简单：使用简单，建立分布式应用更容易。
	   2. 高效：调用过程看起来十分清晰，效率高。
	   3. 通用：进程间通讯的方式，有通用的规则。
	   PRC 调用过程：
	   1. 调用方 client 要使用右侧 server 的功能（方法），发起对方法的调用
	   2.client stub 是 PRC 中定义的存根，看做是 client 的助手。 stub 把要调用的方法参数进行序列化，方法名称和其他数据包装起来。
	   3. 通过网络 socket( 网络通信的技术 ) ，把方法调用的细节内容发送给右侧的 server
	   4.server 端通过 socket 接收请求的方法名称，参数等数据，传给 stub 。
	   5.server 端接到的数据由 serverstub(server 的助手 ) 处理，调用 server 的真正方法，处理业务
	   6.server 方法处理完业务，把处理的结果对象（ Object ）交给了助手，助手把 Object 进行序列化，对象转为二进制数据。
	   7. server 助手二进制数据交给网络处理程序
	   8. 通过网络将二进制数据，发送给 client 。
	   9.client 接数据，交给 client 助手。
	   10.client 助手，接收数据通过反序列化为 java 对象（ Object ），作为远程方法调用结果。

	   rpc 通讯是基于 tcp 或 udp 议
	   序列化方式（ xml/json/ 二进制）


	   Go 的 RPC 只支持go写的系统
	   go rpc的函数有特殊要求：
	     函数首字母必须大写
		 必须只有两个参数，第一个参数是接收的参数，第二个参数是返回给客户端的参数（必须是指针类型）
		 函数必须有一个返回值error
		func(t *T) MethodName(argType T1, replayType *T2) error {

		}

	服务端
	  rpc.Regist(new(结构体))
	  rpc.HandleHttp() 借用http协议作为rpc载体
	  net.Listen("tcp",":1234") 创建一个监听器
	  http.Serve(listen,nil)  启动服务

	客户端
	创建client 使用方法rpc.DialHttp(*协议, 服务名称,地址)
	client.Go(结构体名,方法名,入参,回参指针,chan或nil)
	  返回一个chan 自行创建阻塞时间

	client.Call(结构体名,入参,回参指针)
	直接阻塞
*/
import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
}

type Req struct {
	NumOne int
	NumTwo int
}

type Res struct {
	Num  int
	Code int
}

func (s *Server) Add(req Req, res *Res) error {
	// 异步sleep
	time.Sleep(5 * time.Second)
	res.Num = req.NumOne + req.NumTwo
	if res.Num > 0 {
		res.Code = 200
	}
	return nil
}

func main() {
	// 注册
	rpc.Register(new(Server))
	// 创建Handle
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("error:")
	}
	http.Serve(listen, nil)
}
