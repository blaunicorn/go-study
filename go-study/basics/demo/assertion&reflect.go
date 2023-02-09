package demo

import (
	"fmt"
	"reflect"
)

// demo\assertion&reflect.go

func Assertion() {
	user := User{
		Name: "xiaoming",
		Age:  20,
		Sex:  true,
	}
	check(user)
	user1 := Student{User: User{}}
	user1.Class = "first"
	check(user1)

	fmt.Println("reflect:")
	// reflect.ValueOf() 获取输入参数接口中的数据的值
	// reflect.TypeOf()  动态获取输入参数接口中的值的类型
	// reflect.TypeOf().Kind() 判断类型详细信息
	// v.Kind() == reflect.Float64  使用 Type 的 Kind 方法，和基本数据类型进行相等判断。
	// reflect.ValueOf(varname).Elem() 反射获取变量所指向的指针
	// reflect.ValueOf(x).CanSet() Go 语言反射是否可以修改变量语法：
	// reflect.ValueOf(&x).Elem().Set() Go语言反射修改变量
	// reflect.ValueOf(&x).Elem().SetXXX()
	// reflect.ValueOf().Field(i int) StructField 根据索引，返回索引对应的结构体字段的信息。当值不是结构体或索引超界时会引发 pannic
	// 	NumField() int	返回结构体成员字段数量。当类型不是结构体或索引超界时引发 pannic
	// FieldByName(name string) (StructField, bool)	根据给定字符串返回字符串对应的结构体字段的信息。没有找到时 bool 返回 false，当类型不是结构体或索引超界时引发 pannic
	// FieldByIndex(index []int) StructField	多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。没有找到时返回零值。当类型不是结构体或索引超界时引发 pannic
	// FieldByNameFunc( match func(string) bool) (StructField,bool)	根据匹配函数匹配需要的字段。当值不是结构体或索引超界时引发 pannic
	// 	Go语言解析结构体Tag
	// 语法
	// personType := reflect.TypeOf(person)
	// fieldName, isOk := personType.FieldByName("Name")
	// jsobTagVal := fieldName.Tag.Get("json")
	// 	Go语言反射解析结构体字段值
	// 语法
	// personValue := reflect.ValueOf(person)
	// nameValue := personValue.FieldByName("Name").String()
	// 	Go语言反射解析结构体字段值语法
	// personNameValue := reflect.ValueOf(&person.Name)
	// personNameValue.Elem().SetString("haicoder")
	// 	Go语言反射调用结构体方法
	// 语法
	// personValue := reflect.ValueOf(p)
	// infoFunc := personValue.MethodByName("Info")
	// infoFunc.Call([]reflect.Value{})
	reflectCheck(user)
	user2 := Student{user, "class 2"}
	reflectCheck1(&user2)

}

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	User
	Class string
}

func (user User) SayName(name string) {
	fmt.Println("My name is ", name)
}

func check(v interface{}) {
	switch v.(type) {
	case User:
		fmt.Println("I am user")
	case Student:
		class := v.(Student).Class
		fmt.Println("I am student.", "I am in ", class)
	default:
		fmt.Println("no data")
	}
}

// reflect
func reflectCheck(v interface{}) {
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	//  type, value of v
	fmt.Println(t, value)
	for i := 0; i < t.NumField(); i++ {
		// Get property values by location
		fmt.Println(value.Field(i))
	}
	// fmt.Println(value.FieldByIndex([]int{0, 0}))
}

func reflectCheck1(v interface{}) {
	t := reflect.TypeOf(v)
	value := reflect.ValueOf(v)
	//  type, value of v
	fmt.Println(t, value)
	// for i := 0; i < t.NumField(); i++ {
	// 	// Get property values by location
	// 	fmt.Println(value.Field(i))
	// }

	// fmt.Println("name:", value.FieldByIndex([]int{0, 0}))
	// fmt.Println("name:", value.FieldByName("Name"))

	typeKind := t.Kind()
	if typeKind == reflect.Struct {
		fmt.Println("it is Struck")
	}
	if typeKind == reflect.String {
		fmt.Println("it is String")
	}

	// Modify value.need  transmit pointer
	e := value.Elem()
	e.FieldByName("Class").SetString("class 3")
	fmt.Println(t, value)
}
