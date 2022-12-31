package demo

import "fmt"

var A string = "测试" // go中包所写的要给别的地方引用必须大写开头，就是共有变量，小写就是私有变量
// 关键字 变量名 变量类型 = 变量值

// aaa := "另一种声明变量的方式"  // 需要在程序使用

// 注意：在 Go中是不支持默认参数，需要自己设计
func Demo(a string) {
	aaa := "ceshi" //另一种声明变量的方式
	fmt.Println(a)
	fmt.Println(aaa)
}
