package main

import (
	"fmt"
)

func main() {
	// var a string = "10101"
	// b := "ccc"
	// fmt.Println(a + b)
	// fmt.Println("hello World02")

	// num := 100
	// str := strconv.Itoa(num)
	// // # 表示引号
	// fmt.Printf("type:%T value:%#v\n", str, str)

	// str1 := "110"
	// boo1, err := strconv.ParseBool(str1)
	// if err != nil {
	// 	fmt.Printf("str1: %v\n", err)
	// } else {
	// 	fmt.Println(boo1)
	// }
	// str2 := "t"
	// boo2, err := strconv.ParseBool(str2)
	// if err != nil {
	// 	fmt.Printf("str2: %v\n", err)
	// } else {
	// 	fmt.Println(boo2)
	// }

	// var a *int
	// *a = 100
	// fmt.Println(*a)

	// var b map[string]int
	// b["测试"] = 100
	// fmt.Println(b)

	a := [...]int{1, 34, 43, 34, 33, 34, 3}
	fmt.Println(len(a))

}
