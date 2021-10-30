package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a1 [1]byte
	fmt.Println(a1)
	var a = 65
	var b = string(a)       // A
	var c = strconv.Itoa(a) // 文本类型 65
	var d int
	d, _ = strconv.Atoi(c) // 数字65
	println(a, b, c, d)
}
