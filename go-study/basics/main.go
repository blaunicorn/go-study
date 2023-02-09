// main.go
package main

import (
	"fmt"
	// 包的引用： 一个文件夹下面不能出现多个包。但可以有同一个包名的多个.go文件
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

	// 从包中引用变量和函数
	b := demo.A
	fmt.Println(b)
	demo.Demo("测试1")
	demo.Demo(a)

	// 基本数据类型练习
	// demo.Primitive()

	// 引用数据类型练习
	// demo.Reference()

	// 语句复习
	// demo.Control()

	// 函数方法复习
	// demo.FunctionalMethod()

	// goroutine & channel
	// demo.GoRun()

	// demo\assertion&reflect.go
	// demo.Assertion()

	// demo\sync.go
	// demo.SyncClass()

	// demo\io.go
	// demo.IoClass()

	// demo\net.go
	// demo.NetClass()

	// demo\net-clenit.go
	// demo.NetClientClass()

	// demo\http-class.go
	// demo.HttpClass()

	// demo\http-client-class.go
	demo.HttpClientClass()
}
