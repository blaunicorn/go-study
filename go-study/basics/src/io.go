package main

import (
	"fmt"
	"io"
	"os"
)

func fileTest() {
	f, err := os.OpenFile("./test.txt", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	// 按字节读取
	// for {
	// 	b := make([]byte, 10)
	// 	n, err := f.Read(b)
	// 	fmt.Println(b)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	fmt.Println(string(b), n)
	// }

	// isPrefix 缓存读满标记
	// reader := bufio.NewReader(f)
	// for {
	// 	str, err := reader.ReadString('n')
	// 	if err != nil {
	// 		fmt.Println((err.Error()))
	// 		return
	// 	}
	// 	fmt.Println(str)
	// }

	// 读取整个文件
	// b, _ := ioutil.ReadAll(f)
	// fmt.Println(string(b))

	// 读文件夹?
	// d, err := os.ReadDir("./")
	// fmt.Println(d, err)
	// for _, v := range d {
	// 	fmt.Println(v.Info())

	// 	fmt.Println(v.Name())
	// 	fmt.Println(v.IsDir())
	// 	fmt.Println(v.Type())

	// }

	// 插入复写
	// writer := bufio.NewWriter(f)
	// reader := bufio.NewReader(f)
	// n := 0
	// for {
	// 	n++
	// 	str, err := reader.ReadString('\n')
	// 	writer.WriteString(strconv.Itoa(n) + " " + str)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		break

	// 	}

	// }
	// f.Seek(0, 0)
	// writer.Flush()

	// 复制文件操作
	f2, _ := os.OpenFile("./testCopy.txt", os.O_CREATE|os.O_RDWR, 0777)
	io.Copy(f2, f)
}
