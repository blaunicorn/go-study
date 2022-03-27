// json.go
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

// Golang读取JSON文件
type Tcp_config struct {
	Type          string `json:"type"`
	Ip            string `json:"ip"`
	Port          int    `json:"port"`
	Time_interval int    `json:"time_interval(s)"`
}

func main() {
	// 打开json文件
	josnFile, err := os.Open("tcp_config.json")
	// 最好要处理以下错误
	if err != nil {
		fmt.Println(err)
	}
	// 要预先关闭
	defer josnFile.Close()
	byteValue, _ := ioutil.ReadAll(josnFile)
	fmt.Println(string(byteValue))

	var new_tcp_config Tcp_config
	json.Unmarshal([]byte(byteValue), &new_tcp_config)
	// fmt.Println(new_tcp_config)

	// 如果你很懒不想定义struct，也可以直接使用map[string]interface{}
	// var result map[string]interface{}
	// json.Unmarshal([]byte(byteValue), &result)

	// fmt.Println(result)
	fmt.Println("start the client...")
	ClientBase(&new_tcp_config)
	for range time.Tick(time.Duration(new_tcp_config.Time_interval) * 1000 * time.Millisecond) {
		ClientBase(&new_tcp_config)
	}
}

func ClientBase(new_tcp_config *Tcp_config) {
	port := strconv.Itoa(new_tcp_config.Port)

	ip := new_tcp_config.Ip + ":" + port
	// open connection:
	conn, err := net.Dial("tcp", ip)
	// defer conn.Close()
	if err != nil {
		fmt.Println("error dial:", err.Error())
		return
	}
	// inputContent := "aaaa"
	// str := ip + " say " + inputContent
	currentTime := time.Now()
	// str := "DTUSEND:1#YJ,001234,0010.2,000231,001025,BEIER0,MC2345,2006.01.02 15:04:05"
	// n, err := conn.Write([]byte(inputContent))
	wellArray := []string{"1号测试井", "2号测试井"}
	rand.Seed(time.Now().Unix())
	weight := strconv.Itoa(rand.Intn(1000) + 1000)
	measurement_time := DecimalRoundFix2Str(float64(rand.Intn(300)+100) / 10)
	initial_height := strconv.Itoa((rand.Intn(10) + 10))
	temp_height, err := strconv.Atoi(initial_height)
	termiation_height := strconv.Itoa((rand.Intn(10) + temp_height))
	operatorArray := []string{"张三", "李四", "王五", "赵六"}
	carArray := []string{"班组101", "班组202", "班组303", "班组404"}
	str := "DTUSEND:" + wellArray[rand.Intn(len(wellArray))] + "," + weight + "," + measurement_time + "," + initial_height + "," + termiation_height + "," + operatorArray[rand.Intn(len(operatorArray))] + "," + carArray[rand.Intn(len(carArray))] + "," + currentTime.Format("2006.01.02 15:04:05")
	n, err := conn.Write([]byte(str))
	if err != nil {
		fmt.Println("error write:", err.Error())
		fmt.Println(n)
		return
	} else {
		fmt.Printf("send data from client to ip %s:%s\n", ip, str)

	}
	// defer conn.Close()
}

func DecimalRoundFix2Str(f float64) string {
	f1 := math.Trunc(f*1e2+0.5) * 1e-2
	value := strconv.FormatFloat(f1, 'f', 2, 64)
	return value
}
