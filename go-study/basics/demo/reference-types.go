// demo\reference-types.go
package demo

import (
	"fmt"
)

// 引用类型
//
//	 数组
//		结构 struct
//		接口 interface
//		数组 [length] value type(value1,value2)
//		切片 slice []value type{value1,value2}
//		map  [key type] [value type] [key:value]
//		指针* &
//		函数func
//		管道chan
func Reference() {
	// 数组
	// [element length]element type {element1,element2,...}
	Arr := [3]int{1, 2, 3}                     // fixed length
	c := [...]int{1, 2, 3, 3, 43, 5, 334, 433} // unfixed length
	fmt.Println(Arr, c)
	p := new([10]int) // first declare the array, then assign a value
	p[2] = 5
	fmt.Println(*p, p) //  point to memory address
	animal := [...]string{"dog", "cat", "monkey"}

	for i := 0; i < len(animal); i++ {
		fmt.Print(i)
		fmt.Println(" " + animal[i] + " is crying!")
	}

	// display subscripts and elements
	for k, v := range animal {
		fmt.Println(k, v)
	}

	// append array
	// way 1
	var b string
	b = "test"
	a := []string{"a", "b"}
	a = append(a, b)
	fmt.Println(a)
	// way 2
	arr3 := []string{"test1", "test2"}
	arr2 := []string{"a", "b"}
	arr2 = append(arr3, arr2...)
	fmt.Println(arr2)

	// cap(animal)  capacity  len(animal) length
	fmt.Println(cap(animal), len(animal))

	// two dimensional arry , multldimensional array
	arrays := [3][3]int{{0, 1, 2}, {1, 2, 3}, {3, 4, 5}}
	fmt.Println(arrays)

	// slice
	arr1 := [6]int{0, 1, 2, 3, 4, 5}
	cl := arr1[1:3] // principle of front closing and rear opening
	cl[0] = 5       // slices are part of an array
	fmt.Println(cl)
	fmt.Println(arr1)  // slice changed and array changed
	cl = append(cl, 8) // Ten length and capacity of the array are completely fixed,but the slice can be increased.
	fmt.Println(cl)
	cl[1] = 6
	fmt.Println(cl, arr1) // slice changed and array changed
	cl = append(cl, 9)
	fmt.Println(cl, len(cl), cap(cl)) // slices's capacity and lenght is unfixed.
	cl2 := arr1[2:]
	copy(cl, cl2) //  this is  silces copy
	fmt.Println(cl, cl2)
	// var name []Type
	var aa []int // this  is a slice
	// make( []Type, size, cap )
	aaa := make([]int, 5) // this  is a other slice
	// aaa = append(aa, 10)
	fmt.Println(aa, aaa)

	// map
	fmt.Println("map：")
	// var mapName map[keyType]valueType // declare
	// var m map[string]string
	// m = map[string]string{}  //
	// m1 := map[string]string{}  // Implicit declaration
	m2 := make(map[string]string) // make
	m2["name"] = "wcy"            // assignment
	m2["sex"] = "gentleman"
	fmt.Println(m2, len(m2)) // length

	m3 := map[int]bool{}
	m3[1] = true
	m3[2] = false
	m3[3] = false
	fmt.Println("m3:", m3)

	m4 := map[int]interface{}{} // you can use an empty interface as valueType
	m4[2] = false
	m4[1] = "str-ok"
	m4[3] = 1
	m4[4] = [...]int{1, 2}
	fmt.Println("m4:", m4)

	m5 := map[interface{}]interface{}{} // you can use an empty interface as keyType
	m5["name"] = false
	m5["sex"] = "str"
	m5[3] = 1
	m5[4] = [...]int{5, 6}
	fmt.Println("m5:", m5, len(m5))
	// loop key and value
	for k, v := range m5 {
		fmt.Println(k, v)
	}
	// assignment
	m5[3] = 8

	// delete key
	delete(m5, "sex")
	fmt.Println("deleted m5:", m5)

	// pointer
	fmt.Println("pointer：")
	var p_test int = 123
	fmt.Println(p_test)
	var p_test_b *int = &p_test // this is a pointer type, is Memory address
	fmt.Println(p_test, p_test_b)
	*p_test_b = 999 // Assign the value of the variable pointed to by the pointer
	fmt.Println(p_test, p_test_b)
	fmt.Println(p_test == *p_test_b, &p_test == p_test_b)

	// Define pointer variable
	var house = "Malibu Point 10880, 90265"
	var pointer1 *string = &house
	value := *pointer1 //Value operation on pointer
	*pointer1 = "100"  // Assign the value of the variable pointed to by the pointer
	fmt.Println(house, *pointer1, pointer1, value)

	// array pointer and pointer array
	arr := [5]string{"1", "2", "3", "4", "5"}
	var arrP *[5]string // this is a array pointer
	arrP = &arr
	fmt.Println(arr, arrP, *arrP)
	arrP1 := &arr[1]
	fmt.Println(arrP1, &arr[2], arr[2])

	var arrpp [5]*string // this is a pointer array
	var str1 = "str1"
	var str2 = "str2"
	var str3 = "str3"
	var str4 = "str4"
	var str5 = "str5"
	arrpp = [5]*string{&str1, &str2, &str3, &str4, &str5}
	fmt.Println(arrpp)
	*arrpp[2] = "333"
	fmt.Println(str3)
	// It is recommended to use pointers on methods
	// (provided that this type is not a reference type such as map, slice, channel, etc.)
	// If you want to modify the data or state inside the structure, you must use a pointer
	var str8 = "I am a string"
	fmt.Println(str8)
	pointFunc(&str8)
	fmt.Println(str8)

	// struct
	// this is a data types that can store different types of data
	fmt.Println("struct:")
	//  Explicit declaration
	var xiaoming Man
	xiaoming.Age = 16
	xiaoming.Name = "xiaoming"
	xiaoming.Sex = true
	xiaoming.Hobbies = []string{"game", "sing"}
	xiaoming.Address.City = "daqing"
	fmt.Println(xiaoming)
	fmt.Printf("%#v", xiaoming)

	// implicit declaration
	xiaogao := Man{
		"xiaogao", 18, true, []string{"game", "write"}, Address{"345", "123@124.com"},
	}
	xiaozhang := Man{
		Name: "xiaozhang", Age: 19, Sex: true, Hobbies: []string{"dance", "sing"}, Address: Address{
			Province: "heilongjiang",
			City:     "daqing",
		},
	}
	fmt.Println(xiaogao)
	fmt.Println(xiaozhang)
	// new declaration
	liu := new(Man)
	liu.Name = "liu"
	fmt.Println(liu, liu.Name)
	funcForMan(&xiaozhang)
	re := liu.Sing("the moon")
	fmt.Println((re))
	// struct call its function
	xiaozhang.Run()

	// interface
	fmt.Println("interface:")
	/*
		 	type InterfaceNamer interface{
			     func1(param_list1) return_list1
			     func2(param_list2) return_list2
			    ...
			 }
	*/
	// way 1
	var animal1 Animal
	cat := Cat{
		Name: "tom",
		Sex:  false,
	}
	// Call the method of the original instance with the interface
	animal1 = cat
	animal1.Eat()
	animal1.Run()
	// way 2
	var animal2 Animal
	animal2 = Cat{
		Name: "jack",
		Sex:  true,
	}
	animal2.Eat()
	animal2.Run()

	// Interfaces can be polymorphic
	MyFunc([]string{"1212", "2323"}) // MyFunc can receive any parameter.
	MyFunc(1)
	MyFunc("1")

	// way 3
	cat1 := Cat{
		Name: "wang",
		Sex:  true,
	}
	MyFuncA(cat1)
}

func pointFunc(p1 *string) {
	*p1 = "I sure," + *p1
}

type Man struct {
	Name    string
	Age     int
	Sex     bool
	Hobbies []string
	Address Address
}
type Address struct {
	Province string
	City     string
}

func funcForMan(man *Man) {
	fmt.Println(*man)
}
func (man *Man) Run() {
	fmt.Println(man.Name, " run to ", man.Address.Province, man.Address.City)
}

// the function of the struct
func (man *Man) Sing(name string) (quality string) {
	quality = "very good "
	fmt.Printf("%s sings a %v,people think it is %v", man.Name, name, quality)
	fmt.Println()
	return quality
}

type Animal interface {
	Eat()
	Run()
}
type Cat struct {
	Name string
	Sex  bool
}
type Dog struct {
	Name string
}

func (c Cat) Run() {
	fmt.Println(c.Name, " runs!")
}
func (c Cat) Eat() {
	fmt.Println(c.Name, " eats!")
}

func (d Dog) Run() {
	fmt.Println(d.Name, " runs!")
}
func (d Dog) Eat() {
	fmt.Println(d.Name, " eats!")
}

func MyFunc(a interface{}) {
	fmt.Println(a)
}

func MyFuncA(a Animal) {
	fmt.Println(a)
	a.Run()
	a.Eat()
}
