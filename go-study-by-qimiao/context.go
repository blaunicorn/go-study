package main

import (
	"context"
	"fmt"
	"time"
)

func Context() {
	ctx := context.WithValue(context.Background(), "name", "qm")
	// 通过ctx 把 value值传给子进程
	ctx, clear := context.WithCancel(ctx)
	// flag := make(chan bool)
	message := make(chan int)
	go son(ctx, message)
	for i := 0; i < 10; i++ {
		message <- i
	}
	clear()
	// 结束进程后，再sleep1秒，
	time.Sleep(time.Second)
	fmt.Println("主进程结束")
}

func son(ctx context.Context, msg chan int) {
	t := time.Tick(time.Second)
	for _ = range t {
		select {
		case m := <-msg:
			fmt.Printf("receive the value,%d\n", m)
		case <-ctx.Done():
			fmt.Println("I'm end.")
			fmt.Println(ctx.Value("name")) // 这将显示qm
			return
		}

	}
}

// 原始
// func main() {
// 	flag := make(chan bool)
// 	message := make(chan int)
// 	go son(flag, message)
// 	for i := 0; i < 10; i++ {
// 		message <- i
// 	}
// 	flag <- true
// 	// 结束进程后，再sleep1秒，
// 	time.Sleep(time.Second)
// 	fmt.Println("主进程结束")
// }

// func son(flag chan bool, msg chan int) {
// 	t := time.Tick(time.Second)
// 	for _ = range t {
// 		select {
// 		case m := <-msg:
// 			fmt.Printf("receive the value,%d\n", m)
// 		case <-flag:
// 			fmt.Println("I'm end.")
// 			return
// 		}

// 	}
// }
