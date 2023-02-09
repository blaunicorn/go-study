package demo

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func HttpClientClass() {
	fmt.Println("http client:")

	client := new(http.Client)
	// req, _ := http.NewRequest("GET", "https://www.baidu.com", nil)
	req, _ := http.NewRequest("POST", "http://localhost:8080/test", bytes.NewBuffer([]byte("{\"test\":\"这是客户端\"}")))
	// 也可以构建header
	req.Header["test"] = []string{"test1 test2"}
	res, _ := client.Do(req)
	body := res.Body
	b, _ := io.ReadAll(body)
	fmt.Println(string(b))
}
