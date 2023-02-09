package main

import (
	"fmt"
	"log"
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
	Code int
	Num  int
}

func main() {
	req := Req{NumOne: 6, NumTwo: 5}
	var res Res
	client, err := rpc.DialHTTP("tcp", "localhost:8888")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// client.Call("Server.Add", req, &res) // 同步
	// fmt.Println(res)

	call := client.Go("Server.Add", req, &res, nil) //异步

	fmt.Println("异步")
	// <-call.Done
	for {
		select {
		case <-call.Done:
			fmt.Println(res)
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("waiting...")
		}
	}

}
