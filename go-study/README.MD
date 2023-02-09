# 通过bilibili里的教学视频复习go的基础知识
### 下载go的安装包
go官网：https://golang.google.cn/
go语言中文网：https://studygolang.com/
## 采用vscode作为工具
采用vscode开发go
安装插件go

## 第一个程序
新建main.go文件。并编辑，测试是否可以运行
```go
// main.go
package main

import "fmt"

// 主入口
func main() {
	fmt.Println("hello main!!")
}

```
在命令行上运行
```cmd
go  run main.go  
```

### go的一般目录结构
├──go-demo
    ├──bin  // 可执行文件
    ├──pkg  // 包
    ├──src  //项目
    ├── go.mod  // 模块列表管理
    └──main.go   // main包 main方法

###   main入口文件和包的引用
```go
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
}

```
 新建demo文件夹，并建立my-string.go文件
```go
// demo\my-string.go
package demo

import "fmt"

var A string = "变量，首字母大写，可以被别的包引用" // go中包所写的要给别的地方引用必须大写开头，就是共有变量，小写就是私有变量
var a string = "首字母小写，私有变量，不能别其他包使用"

// 关键字 变量名 变量类型 = 变量值

// aaa := "另一种声明变量的方式"  // 需要在程序使用

// 注意：在 Go中是不支持默认参数，需要自己设计
func Demo(a string) {
	aaa := "ceshi" //另一种声明变量的方式
	fmt.Println(a)
	fmt.Println(aaa)
}

//  双斜杠为单行注释

/*
   区块注释
*/

```