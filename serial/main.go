// project serial-demo

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"time"

// 	serial "github.com/tarm/goserial"
// )

// const MAXRWLEN = 8000

// func main() {

// 	cfg := &serial.Config{Name: "COM1", Baud: 9600, ReadTimeout: 3 /*毫秒*/}
// 	fmt.Println("【打开端口中...】")
// 	iorwc, err := serial.OpenPort(cfg)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	defer iorwc.Close()
// 	buffer := make([]byte, MAXRWLEN)

// 	//发命令之前清空缓冲区
// 	num, err := iorwc.Read(buffer)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	//发命令数据类型为[]byte
// 	num, err = iorwc.Write([]byte("AT\r\n"))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	var tmpstr string = ""
// 	for i := 0; i < 3000; i++ {
// 		fmt.Println("【INFO】采集数据中...")
// 		num, err = iorwc.Read(buffer)
// 		if num > 0 {
// 			tmpstr += fmt.Sprintf("%s", string(buffer[:num]))
// 			fmt.Println(tmpstr)
// 		}

// 		//查找读到信息的结尾标志
// 		if strings.LastIndex(tmpstr, "\r\nOK\r\n") > 0 {
// 			break
// 		}

// 		fmt.Println("【INFO】数据更新成功！")
// 		// 线程休眠 5 min
// 		time.Sleep(time.Duration(100) * time.Second)
// 	}

// 	//打印输出读到的信息
// 	fmt.Println(tmpstr)

// 	return
// }

package main

import (
	"flag"
	"log"
	"os"

	"github.com/larspensjo/config"
	serial "github.com/tarm/goserial"
)

var (
	conFile = flag.String("configfile", "/config.ini", "config file")
)

var TOPIC = make(map[string]string)

func main() {
	//获取当前路径
	file, _ := os.Getwd()
	cfg, err := config.ReadDefault(file + *conFile)

	//获取配置文件中的配置项
	id, err := cfg.String("COM", "COMID")
	//设置串口编号
	c := &serial.Config{Name: id, Baud: 115200}
	//打开串口
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	command, err := cfg.String("COM", "COMMAND")
	// 写入串口命令
	log.Printf("写入的指令 %s", command)
	n, err := s.Write([]byte(command))

	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 128)
	n, err = s.Read(buf)
	log.Printf("读取窗口信息 %s", buf[:n])
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%q", buf[:n])
}
