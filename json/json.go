// json.go
package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name string
	Age  int
}

type Class struct {
	Students []Student
}

// func main() {
// 	var s Class
// 	str := `{"Students":[{"Name":"zs","Age":13},{"Name":"ls","Age":"2"}]}`
// 	json.Unmarshal([]byte(str), &s)
// 	fmt.Println(s)
// 	fmt.Println(s.Students[0].Name)

// 	fmt.Println(json.Marshal(s))
// 	b, _ := json.Marshal(s)
// 	fmt.Println("b:", b)

// 	os.Stdout.Write(b)
// }

// 总结

// 3.1如果使用整形，赋值字符串时，Unmarshar将是0

// 3.2直接Println和取值再Println直接输出对象显示的都是ASCII码

// 3.3使用Stdout.Write可以直接输出对象

// Golang读取JSON文件
// type User struct {
// 	FirstName string `json:"first_name"`
// 	LastName  string `json:"last_name"`
// 	BlogURL   string `json:"blog_url"`
// 	BlogName  string `json:"blog_name"`
// }

// func main() {
// 	// 打开json文件
// 	josnFile, err := os.Open("user.json")
// 	// 最好要处理以下错误
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// 要预先关闭
// 	defer josnFile.Close()
// 	byteValue, _ := ioutil.ReadAll(josnFile)
// 	fmt.Println(string(byteValue))

// 	var user User
// 	json.Unmarshal([]byte(byteValue), &user)
// 	fmt.Println(user)
// 	fmt.Println(user.BlogName)

// 	// 如果你很懒不想定义struct，也可以直接使用map[string]interface{}
// 	var result map[string]interface{}
// 	json.Unmarshal([]byte(byteValue), &result)

// 	fmt.Println(result)
// }

// 实现向管道里写100个数，再读出
// func readCh(ch1 chan int) {
// 	for n := 0; n < 10; n++ {
// 		m := <-ch1
// 		println(m)
// 	}
// }
// func writeCh(ch1 chan int) {
// 	for n := 0; n < 10; n++ {
// 		ch1 <- n
// 	}
// 	println("write end.", &ch1)
// }
// func main() {
// 	ch1 := make(chan int, 100)
// 	writeCh(ch1)
// 	readCh(ch1)
// }

// MakeTimestamp returns the Unix timestamp in milliseconds
func MakeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// 什么是Unix时间戳(Unix timestamp)： Unix时间戳(Unix timestamp)，或称Unix时间(Unix time)、POSIX时间(POSIX time)，是一种时间表示方式，定义为从格林威治时间1970年01月01日00时00分00秒起至现在的总秒数。Unix时间戳不仅被使用在Unix系统、类Unix系统中，也在许多其他操作系统中被广泛采用
func main() {
	i := MakeTimestamp()
	fmt.Println(i)
	// return
}
