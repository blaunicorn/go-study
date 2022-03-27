package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	//初始化测试服务器
	handler := http.HandlerFunc(Index)
	app := httptest.NewServer(handler)
	defer app.Close()
	//测试代码
	//发送http.Get请求，获取请求结果response
	response, _ := http.Get(app.URL + "/index")
	//关闭response.Body
	defer response.Body.Close()
	//读取response.Body内容，返回字节集内容
	bytes, _ := ioutil.ReadAll(response.Body)
	//将返回的字节集内容通过string()转换成字符串，并显示到日志当中
	t.Log(string(bytes))
}

// go test
