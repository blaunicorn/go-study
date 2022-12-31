通过视频复习go的基础知识
go官网：https://golang.google.cn/
go语言中文网：https://studygolang.com/
采用vscode开发go
安装插件go
### 一、基本语法

#### 第一个程序
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
./fresh  run main.go
```

#### go的一般目录结构
├──go-demo
    ├──bin  // 可执行文件
    ├──pkg  // 包
    ├──src  //项目
    ├── go.mod  // 模块列表管理
    └──main.go   // main包 main方法
