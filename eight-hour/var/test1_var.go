// 四种变量的声明方式
package main

import "fmt"

func main() {
	// 方法一 声明变量，默认的值为0/
	var a int
	fmt.Println(a)

	// 方法二 声明一个变量，初始化一个值const
	var b int = 100
	fmt.Println(b)
	fmt.Printf("type of b=%T\n", b)
	var bb string = "abcd"
	fmt.Printf("bb=%s,type of bb=%T\n", bb, bb)

	// 方法三，初始化时，省去数据类型，通过值自动匹配推断数据类型
	var c = 200
	fmt.Println("c=", c)
	fmt.Printf("c=%d,type of c=%T\n", c, c)

	// 方法四 省略var关键字，直接匹配
	d := 100
	fmt.Printf("d=%d,type of d=%T\n", d, d)
	e := 1.24
	fmt.Printf("d=%f,type of d=%T\n", e, e)
}
