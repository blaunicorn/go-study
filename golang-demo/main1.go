package main

import "fmt"

// class1
// import (
// 	"fmt"
// 	module_test "test/module"
// )

// func main() {
// 	var a string = "hello aaaa"
// 	fmt.Println(a)
// 	fmt.Println(module_test.A)
// 	fmt.Println(module_test.B)
// }

func main1() {
	var num1 uint = 999
	var num2 int = -999
	var num3 float64 = 3.141592664
	var string1 string = "3.145555"
	var bool1 bool = true
	fmt.Printf("%T", num1) // 当前数据类型
	fmt.Println(num1, num2, num3, string1, bool1)
}
