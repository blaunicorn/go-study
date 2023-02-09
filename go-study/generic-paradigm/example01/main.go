package main

// 泛型01: 只有在调用时 才能确定类型
/*
跟声明接口一样，使用type x interface{} 关键字来声明，不过里面的成员不再是方法，而是类型，类型之间用符号 "|" 隔开
type MyInt interface {
	int | int8 | int16 | int32 | int64
}

any: 表示go里面所有的内置基本类型，等价于interface{}
comparable: 表示go里面所有内置的可比较类型：int、uint、float、bool、struct、指针等一切可以比较的类型
符号"~"都是与类型一起出现的，用来表示支持该类型的衍生类型


// 使用泛型
func GetMaxNum[T int | int8](a, b T) T {
	if a > b {
		return a
	}

	return b
}

// 像声明接口一样声明
type MyInt interface {
	int | int8 | int16 | int32 | int64
}

// T的类型为声明的MyInt
func GetMaxNum[T MyInt](a, b T) T {
	if a > b {
		return a
	}

	return b
}


创建一个带有泛型的结构体User，提供两个获取age和name的方法
注意：只有在结构体上声明了泛型，结构体方法中才可以使用泛型

type AgeT interface {
	int8 | int16
}

type NameE interface {
	string
}

type User[T AgeT, E NameE] struct {
	age  T
	name E
}

// 获取age
func (u *User[T, E]) GetAge() T {
	return u.age
}


// 获取name
func (u *User[T, E]) GetName() E {
	return u.name
}

// 声明要使用的泛型的类型
var u User[int8, string]

// 赋值
u.age = 18
u.name = "weiwei"

// 调用方法
age := u.GetAge()
name := u.GetName()

// 输出结果 18 weiwei
fmt.Println(age, name)

*/

import (
	"fmt"
	"strconv"
)

func conv(i int) string {
	return strconv.Itoa(i)
}

// 泛型
func ConvGeneric[T any](i T) T {
	return i
}

// 定义泛型结构体
type User[A any, B any] struct {
	Name A
	Age  B
}

type Slice1 []string        // 正常的切片
type CustomSlice[A any] []A // 定义泛型切片, 这样切片的类型就是可变的,

// 定义泛型map, 需要进行类型约束 type CustomMap[K any, V any] map[K]V 这样没有约束不行
type CustomMapKey interface {
	string | int | bool
}

type CustomMap[K CustomMapKey, V any] map[K]V

// 用方法约束泛型
type MyType interface {
	getValue() string
}

func test1[T MyType](t T) {
	fmt.Println(t.getValue())
}

type my struct {
	Name string
}

func (m my) getValue() string {
	return m.Name
}
func main() {
	fmt.Println("类型转换示例：", conv(123))
	// 同样方法，传入不同类型，得到不同的结果。这样就减少了代码量
	fmt.Println("使用泛型做类型转换string：", ConvGeneric("123")+"345")
	fmt.Println("使用泛型做类型转换int：", ConvGeneric(123)+345)

	// 泛型用于结构体
	user := User[string, int]{
		Name: "zhagnsan",
		Age:  20,
	}
	fmt.Println(user)

	// 泛型用于切片
	customSlice1 := CustomSlice[int]{1, 2, 3, 4}
	customSlice2 := CustomSlice[string]{"a1", "a2", "a3", "a4"}
	fmt.Println(customSlice1, customSlice2)

	customMap := make(CustomMap[string, int])
	customMap["age"] = 18
	fmt.Println(customMap)
}
