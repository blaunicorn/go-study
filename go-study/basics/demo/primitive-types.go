// demo\primitive-types.go
package demo

// go的基本类型
import (
	"fmt"
	"strconv"
)

func Primitive() {
	fmt.Println("基本类型:")

	// 基本类型和变量
	// 整型 int int8 int16(-32768-32767) int32 int64 uint uint8(0-255) uint16(0-65535) uint32 uint64 byte(字节存储) rune(等价于int32 存储一个unicode编码)
	// int类型的大小为 8 字节;int8类型大小为 1 字节(-128-127);int16类型大小为 2 字节;int32类型大小为 4 字节;int64类型大小为 8 字节
	// 浮点型 float 32 64
	// 字符串 string
	// bool true false
	// 注意 双引号 单引号 反引号的区别
	// 	双引号用来创建可解析的字符串字面量(支持转义，但不能用来引用多行)
	// 反引号用来创建原生的字符串字面量，这些字符串可能由多行组成(不支持任何转义序列)，原生的字符串字面量多用于书写多行消息、HTML以及正则表达式
	// 而单引号则用于表示Golang的一个特殊类型：rune，类似其他语言的byte但又不完全一样，是指：码点字面量（Unicode code point），不做任何转义的原始内容。
	fmt.Println("基本类型之整型:")
	var num0 uint = 999
	var num1 int = -999
	fmt.Printf("整型:%d;%d.\n", num0, num1)

	fmt.Println("基本类型之浮点型:")
	var num3 float32 = 3.1415926
	fmt.Printf("%f\n", num3)

	fmt.Println("基本类型之布尔型:")
	var bool1 bool = true
	var bool2 bool = false
	fmt.Printf("%+v，%v\n", bool1, bool2)

	// ps:“%+v”会以字段键值对的形式key-value格式打印，“%v”只会打印字段值value信息
	type user struct {
		Name string
		Age  int
	}
	userInfo := user{
		Name: "Bill",
		Age:  25,
	}
	// 结构体打印(json格式等...)
	fmt.Printf("type:%T,%+v\n", userInfo, userInfo) // {Name:Bill Age:25}
	fmt.Printf("%v\n", userInfo)

	// 数据类型转换
	// strconv.Atoi(string) 字符串转整型 strconv.ParsInt(string,10,64) 字符串转浮点
	// strcov.Itoa(int) 整型转字符串 strconv.FormatInt(int64,10) int64转string
	// ParseFloat(string,32) 字符串转float32、float64   float32,err = ParseFloat(string,32) float64,err=ParseFloat(sting,64)
	// int64 to int  int:=int(int64)
	// int to int64 int64:= int64(int)
	var str string = "123"
	int1, err := strconv.Atoi(str)
	fmt.Printf("%v,%T：%d,%T,%v\n", str, str, int1, int1, err)

	num2 := 100
	str2 := strconv.Itoa(num2)
	// #v表示用双引号包裹变量(生成该值的源代码片段) t表示变量类型 s表示字符类型变量 d表示数字变量
	fmt.Printf("type:%T value:%#v  value:%s\n", str2, str2, str2)

	// bool
	str1 := "110"
	boo1, err := strconv.ParseBool(str1)
	if err != nil {
		fmt.Printf("str1: %v\n", err)
	} else {
		fmt.Println(boo1)
	}
	str3 := "t"
	boo2, err := strconv.ParseBool(str3)
	if err != nil {
		fmt.Printf("str2: %v\n", err)
	} else {
		fmt.Println(boo2)
	}
}
