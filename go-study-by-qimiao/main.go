// main.go
package main

import (
	"fmt"
	// 需要从go.mod文件作为绝对路径import。并且路径开头应该是module名，不是目录名
	// 此文中是go-study-demo
	// "go-study-demo/testpackage"  // 直接引用
	// root "go-study-demo/testpackage"  // 别名引入
	// . "go-study-demo/testpackage" // 隐式引用
	"go-study-demo/demo"
)

// 主入口
func main() {
	fmt.Println("hello main!!")

	var a string = "bianliang"
	fmt.Printf("变量是:%s", a)
	fmt.Println()

	a = demo.A
	fmt.Println(a)

	demo.Demo("测试1")
}
