package main

import (
	// 需要从go.mod文件作为绝对路径import。并且路径开头应该是module名，不是目录名
	// 此文中是blog
	// "blog/testpackage"  // 直接引用
	// root "blog/testpackage"  // 别名引入

	. "blog/testpackage" // 隐式引用
	"fmt"
	"sync"
)

func main() {
	// 显式变量声明
	// var a string = "10101"
	// 隐式变量声明
	// b := "ccc"
	// fmt.Println(a + b)
	// fmt.Println("hello World02")
	// fmt.Println(testpackage.A) // 包的直接引用
	// fmt.Println(root.A)  // 包的别名引用

	// fmt.Println(A) // 包的隐式引用
	fmt.Println(B) // 包的隐式引用
	fmt.Println(C) // 包的隐式引用
	// num := 100
	// str := strconv.Itoa(num)
	// // # 表示双引号 t表示类型 s表示字符类型 d表示数字
	// fmt.Printf("type:%T value:%#v  value:%s\n", str, str, str)

	// 20210916
	// 基本类型和变量 整形 int 8 16 32 64 uint 8 16 32 64 浮点型 float 32 64  字符串 string  bool true false
	// 注意 双引号 单引号 反引号的区别
	// var num1 uint = 999
	// var num2 int = -999
	// println(num1, num2)
	// var num3 float32 = 3.1415926
	// fmt.Printf("%f", num3)
	// var bool1 bool = true
	// fmt.Printf("%+v", bool1)

	// 数据类型转换  strconv.Atoi 字符串转整形 strconv.ParsInt(string,10,64) 字符串转浮点
	// strcov.Itoa(int)  strconv.FormatInt(int64,10) int64 转string
	// 字符串转 float32 float64  float32,err = ParseFloat(string,32) float64,err=ParseFloat(sting,64)
	// int64 to int  int:=int(int64)
	// int to int64 int64:= int64(int)

	//  结构 struct
	//  接口 interface
	//  数组 [length] value type(value1,value2)
	//  切片 slice []value type{value1,value2}
	//  map  [key type] [value type] [key:value]
	//  指针* &
	//  函数func
	//  管道chan

	// 流程控制语句
	// ++
	// --  no front end
	// a := 0
	// a++
	// fmt.Println(a)
	// 条件语句
	// a := 2
	// if a == 1 {
	// 	fmt.Println(a)
	// } else if a == 2 {
	// 	fmt.Println(a)
	// } else {
	// 	fmt.Println(a)
	// }
	// 选择语句
	// a := 1
	// switch a {
	// case 0:
	// 	fmt.Println("a=0")
	// 	fallthrough // 穿过，可以再走下一个
	// case 1:
	// 	fmt.Println("a=1")
	// default:
	// 	fmt.Print(("都不是"))
	// }

	// 循环语句
	// count := 10
	// for i := 0; i < count; i++ {
	// 	fmt.Println(i)
	// }
	// for {
	// 	count++
	// 	fmt.Println(count)
	// 	if count > 20 {
	// 		break
	// 	}
	// }

	// 跳转语句
	// break 与 continue配合标签可永不多层循环嵌套的跳出const
	// goto 是调整执行位置
	// A:

	// 	for i := 0; i < 10; i++ {
	// 		if i == 1 {
	// 			continue  // 跳出本次循环
	// 		}
	// 		fmt.Println("我是A", i)
	// 		if i > 3 {
	// 			// 跳出A
	// 			break A
	// 			// 跳到B
	// 			goto B
	// 		}
	// 	}

	// B:
	// 	fmt.Println("我是B")
	// 	fmt.Println("结束")

	// 数组
	// [element length]element type {element1,element2,...}
	// Arr := [3]int{1, 2, 3}                     // fixed length
	// c := [...]int{1, 2, 3, 3, 43, 5, 334, 433} // unfixed length
	// fmt.Println(Arr, c)
	// p := new([10]int) // first declare the array and then assign a value
	// p[2] = 5
	// fmt.Println(*p, p) //  point to memory address
	// animal := [...]string{"dog", "cat", "monkey"}

	// for i := 0; i < len(animal); i++ {
	// 	fmt.Println(animal[i] + "is running!")
	// }
	// // cap(animal)  capacity  len(animal) length
	// // display subscripts and elements
	// for k, v := range animal {
	// 	fmt.Println(k, v)
	// }

	// // two dimensional arry , multldimensional array
	// er := [3][3]int{{0, 1, 2}, {1, 2, 3}, {3, 4, 5}}
	// fmt.Println(er)

	// // slice
	// arr1 := [6]int{0, 1, 2, 3, 4, 5}
	// cl := arr1[1:3] // principle of front closing and rear opening
	// cl[0] = 5       // slices are part of an array
	// fmt.Println(cl)
	// fmt.Println(arr1)
	// cl = append(cl, 8) // Ten length and capacity of the array are completely fixed,but the slice can be increased.
	// fmt.Println(cl)
	// cl[1] = 6
	// fmt.Println(cl, arr1)
	// cl = append(cl, 9)
	// fmt.Println(cl, len(cl), cap(cl)) // slices's capacity and lenght is unfixed.
	// cl2 := arr1[2:]
	// copy(cl, cl2) //  this is  silces copy
	// fmt.Println(cl, cl2)
	// var aa []int // this  is a slice
	// aaa := make([]int, 5)
	// // aaa = append(aa, 10)
	// fmt.Println(aa, aaa)

	// map
	// var m map[string]string // declare
	// m = map[string]string{}
	// m1 := map[string]string{}
	// m2 := make(map[string]string)
	// m2["name"] = "qm"
	// m2["sex"] = "man"
	// fmt.Println(m2)

	// m1 := map[int]bool{}
	// m1[1] = true
	// m1[2] = false
	// fmt.Println(m1)
	// m1 := map[int]interface{}{} // you can use an empty interface
	// m1[2] = false
	// m1[1] = "strin"
	// m1[3] = 1
	// m1[4] = [...]int{}
	// fmt.Println(m1)

	// m1 := map[interface{}]interface{}{} // you can use an empty interface
	// m1["name"] = false
	// m1["sex"] = "strin"
	// m1[3] = 1
	// m1[4] = [...]int{}
	// fmt.Println(m1, len(m1))
	// // loop key and value
	// for k, v := range m1 {
	// 	fmt.Println(k, v)
	// }
	// // assignment
	// m1[3] = 8
	// // delete key

	// delete(m1, "sex")
	// fmt.Println(m1)

	// func ()  {

	// }    // note: a function cannot be declared inside a function in Go language,but anonymous functions can be placed.
	//  note: Uppercase can be called by other packages, and lowercase is private to this package.
	// func(input-parameter1 type,input-parameter2 type) (out-parameter1 type,out-parameter2 type) {
	// 	function body
	// }
	// a1(3, "now value >1")
	// a1(0, "now value <1")
	// a2(3, "now value >1")
	// a2(0, "now value <1")
	// r1, r2 := a2(3, "now value <1")
	// // r1, r2 := a2(0, "now value <1")
	// fmt.Println(r1, r2)
	// b := func(data1 string) {
	// 	fmt.Println(data1)
	// }
	// b("I am a anonmous function")

	// mo(9527, "1", "2", "4", "5")
	// ar := []string{"1", "2", "3", "4"}
	// mo(888, ar...)
	// defer firstFunction() //  be defferred ,it will  carry out the final execution.
	// // self executing function
	// func() {
	// 	fmt.Println("I am a self execution function!")
	// }()

	// closeFunction()(4)

	// if a==1 {}
	// if else
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

	// pointer
	// var p_test int
	// p_test = 123
	// fmt.Println(p_test)
	// var p_test_b *int
	// // p_test_b = p_test
	// p_test_b = &p_test
	// fmt.Println(p_test_b)
	// *p_test_b = 999
	// fmt.Println(p_test, p_test_b)
	// fmt.Println(p_test == *p_test_b, &p_test == p_test_b)

	// var a *int
	// *a = 100
	// fmt.Println(*a)

	// array pointer and pointer array

	// arr := [5]string{"1", "2", "3", "4", "5"}
	// var arrP *[5]string
	// arrP = &arr
	// fmt.Println(arr, arrP, *arrP)
	// arrP1 := &arr[1]
	// fmt.Println(arrP1)

	// var arrpp [5]*string
	// var str1 = "str1"
	// var str2 = "str2"
	// var str3 = "str3"
	// var str4 = "str4"
	// var str5 = "str5"
	// arrpp = [5]*string{&str1, &str2, &str3, &str4, &str5}
	// fmt.Println(arrpp)
	// *arrpp[2] = "555"
	// fmt.Println(str3)
	// var b map[string]int
	// b["测试"] = 100
	// fmt.Println(b)
	// var str1 = "我定义了"
	// pointFunc(&str1)
	// fmt.Println(str1)
	// a := [...]int{1, 34, 43, 34, 33, 34, 3}
	// fmt.Println(len(a))
	// var str1 = "我是来测地址的"
	// p := &str1
	// *p = "1212"

	// fmt.Println(str1)
	// var qm Qimiao //  Explicit declaration
	// qm.Age = 18
	// qm.Name = "qimiao"
	// qm.Sex = true
	// qm.Hobbys = []string{"play", "sing"}
	// fmt.Println(qm)

	// implicit declaration
	// qm := Qimiao{
	// 	"qimiao", 18, true, []string{"play", "sing"},
	// }
	// qm := Qimiao{
	// 	Name: "qimiao", Age: 18, Sex: true, Hobbys: []string{"play", "sing"},
	// }
	// fmt.Println(qm)

	// // qm := new(Qimiao)
	// // qm.Name = "haha"
	// // fmt.Println(qm)

	// re := qm.Song("月亮之上")
	// fmt.Println((re))

	// MyFunc([]string{"1212", "2323"})

	var wg sync.WaitGroup
	wg.Add(1)
	go Run(&wg)
	wg.Wait()

}
func Run(wg *sync.WaitGroup) {
	fmt.Println("我跑起来了")
	wg.Done()
}

func MyFunc(a interface{}) {
	fmt.Println(a)
}

type Qimiao struct {
	Name   string
	Age    int
	Sex    bool
	Hobbys []string
}

func (q *Qimiao) Song(name string) (restr string) {
	restr = "ooo ahfa "
	fmt.Printf("%v唱了一首%v,观众觉得%v", q.Name, name, restr)
	fmt.Println()
	return restr
}
func pointFunc(p1 *string) {
	*p1 = "我变脸," + *p1
}
