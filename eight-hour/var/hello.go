package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("hello Go!")
	time.Sleep(5 * time.Second)
	fmt.Println("after is 5s,hello Go!")
}
