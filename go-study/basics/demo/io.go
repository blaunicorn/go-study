package demo

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
)

func IoClass() {
	fmt.Println("io and file:")

	/*
				type Reader interface{
					Read(p []byte) (n int,err error)
				}
				将len(p)个字节读取到p中，ReadFrom() 实现Reader接口

		      type Writer interface{
				Write(p []byte) (n int, err error)
			  }
			  Write方法用于将p中的数据按字节写入到对象的数据流中

			  type Seeker interface{
				Seek(offset int64,whence int) (ret int64, err error)
			  }
			  Seek 设置下一次读写操作的指针位置，每次的读写操作都是从探针位置开始
			  whence 0 从数据的开头开始移动指针，1 从数据的当前指针位置开始移动指针，2 从数据的稳步开始移动指针

			  type Closer interface {
				Close() error
			  }
			  Close 一般用于关闭文件、连接、数据库等

			  写 文件 有四种方法，分别为：使用 io.WriteString 写文件，使用 ioutil.WriteFile 写文件，使用 file.Write 写文件，使用 writer.WriteString 写文件。
	*/

	// os.Open()
	// os.ReadFile()
	file, err := os.OpenFile("./upload/test.txt", os.O_CREATE|os.O_RDWR, 0777) // 调用文件的基础,参数为如果没有创建，并且可读写
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//写入文件
	var content = "hello txt"
	n, err := file.Write([]byte(content))
	if err != nil {
		fmt.Println("Write file err =", err)
		return
	}
	fmt.Println("Write file success, n =", n)
	//读取文件
	// 方式一
	// fileContent, err := ioutil.ReadFile(fileName)
	// 方式二
	// file.Seek(0, 0)
	// for {
	// 	b := make([]byte, 5)
	// 	n, err := file.Read(b)
	// 	if err != nil {
	// 		fmt.Println("read file err =", err)
	// 		return
	// 	}
	// 	fmt.Println("Read file success =", string(b), " n=", n)
	// }

	// fmt.Println("Read file success =", string(fileContent))

	// file.Seek(0,io.SeekEnd) // 从尾部插入
	file.WriteString("\ninsert a line！")

	// 缓存读
	// reader := bufio.NewReader(file)
	// file.Seek(0, 0)
	// for {
	// 	str, err := reader.ReadString('\n') // 以\n标记分割，
	// 	if err != nil {
	// 		fmt.Println((err.Error()))
	// 		return
	// 	}
	// 	fmt.Println(str)
	// }

	// 读取整个文件
	file.Seek(0, 0)
	b, _ := ioutil.ReadAll(file) // ioutil即将废弃,推荐使用os包
	fmt.Println(string(b))

	b1, _ := os.ReadFile("./upload/test.txt")
	fmt.Println("ReadFile:\n", string(b1))

	// 读文件夹
	d, err := os.ReadDir("./upload")
	fmt.Println(d, err)
	for _, v := range d {
		fmt.Println(v.Info())
		fmt.Println("name:", v.Name())
		fmt.Println("isDir:", v.IsDir())
		fmt.Println("type:", v.Type())
	}

	file.Seek(0, 0)
	writer := bufio.NewWriter(file)
	reader := bufio.NewReader(file)
	num := 0
	for {
		num++
		str, err := reader.ReadString('\n')
		writer.WriteString(strconv.Itoa(num) + " " + str) // 写入缓存，还没有写入文件
		if err != nil {
			fmt.Println(err)
			break

		}
	}
	file.Seek(0, 0) // 从头写入，如果不加，则是从尾写入
	writer.Flush()  //   写入文件

	// 复制文件操作
	file.Seek(5, 0) // 指针移到头部，偏移5个字节
	f2, _ := os.OpenFile("./upload/test-copy.txt", os.O_CREATE|os.O_RDWR, 0777)
	io.Copy(f2, file)
}
