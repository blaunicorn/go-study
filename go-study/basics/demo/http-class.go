package demo

import (
	"fmt"
	"io"
	"net/http"
)

func HttpClass() {
	fmt.Println("http:")
	/*

		server 服务端 ，包括地址 端口 处理器等
		conn 链接，用户请求等
		response 响应信息
		request 请求信息
		Handle 对于接收的信息进行处理并返回的处理器

		服务端
		Handle HnadleFunc
		// 创建一个handle
		func handle(res http.ResponseWriter,req *http.Request) {
			fmt.Fprintln(res,"hello world")
		}

		// 霸handle创建，进入默认路由器中
		http.HandleFunc("/handle")

		 NewServerMux // 创建自定义的MUX路由器

		 通过Header()设置头map
		 通过WriteHeader(code) 设置状态码
		 通过Write写入io.writer 来控制返回值
		 通过req.Body  获取请求的body内容

		 http.Handle("/",http.FileServer(http.Dir("/")))  //设置静态显示

		客户端
		创建客户端  client:=new(http.Client)
		构建请求  request,err:=http.newRequest("方法","路径",body)
		发送请求，并得到返回值 res,err:=client.Do(request)

	*/

	// 默认
	// http.HandleFunc("/api/test", handle)
	// http.ListenAndServe(":8080", nil)
	// 浏览器访问 http://localhost:8080/test

	// 独立路由器
	mux := http.NewServeMux()
	mux.HandleFunc("/api/test1", handle)
	http.ListenAndServe(":8081", mux)
}

func handle(res http.ResponseWriter, req *http.Request) {
	// 直接访问，默认为get请求
	// res.Write([]byte("返回的数据"))
	switch req.Method {
	case "GET":
		// 浏览器访问 http://localhost:8080/test
		res.Write([]byte("get请求返回的数据"))
		break
	case "POST":
		res.Write([]byte("post请求返回的数据\n"))
		// post访问 http://localhost:8080/test
		header := res.Header()
		header["test"] = []string{"test1,test2"} // 设置自定义头
		// res.WriteHeader(http.StatusBadRequest)   // 设置自定义状态码
		b, _ := io.ReadAll(req.Body)
		res.Write(b)
		break
	}
}
