package main

import (
	"log"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//ParseFiles从"index.html"中解析模板。
	//如果发生错误，解析停止，返回的*Template为nil。
	//当解析多个文件时，如果文件分布在不同目录中，且具有相同名字的，将以最后一个文件为主。
	files, _ := template.ParseFiles("template/range.html")
	//声明变量b，类型为字符串切片，有三个内容
	b := []string{"张无忌", "张三丰", "张翠山"}
	//声明变量b，类型为字符串切片，没有内容，是为了else示例
	//b := []string{}
	//Execute负责渲染模板，并将b写入模板。

	_ = files.Execute(w, b)
	//w.Write([]byte("/index"))
}

func IndexMap(w http.ResponseWriter, r *http.Request) {
	//ParseFiles从"index.html"中解析模板。
	//如果发生错误，解析停止，返回的*Template为nil。
	//当解析多个文件时，如果文件分布在不同目录中，且具有相同名字的，将以最后一个文件为主。
	files, _ := template.ParseFiles("./template/map.html")
	//声明变量b，类型为map切片，有三个内容
	b := []map[string]interface{}{
		{"id": 1, "name": "张无忌", "age": 18},
		{"id": 2, "name": "周芷若", "age": 16},
		{"id": 3, "name": "谢逊", "age": 39},
	}
	//Execute负责渲染模板，并将b写入模板。
	_ = files.Execute(w, b)
}

// username 需要注意的是，如果采用结构体类型，那么就要考虑字段的首字母必须大写
type Username struct {
	Id   int
	Name string
	Age  int
}

func IndexStruct(w http.ResponseWriter, r *http.Request) {
	//ParseFiles从"index.html"中解析模板。
	//如果发生错误，解析停止，返回的*Template为nil。
	//当解析多个文件时，如果文件分布在不同目录中，且具有相同名字的，将以最后一个文件为主。
	files, _ := template.ParseFiles("./template/struct.html")
	//声明变量b，类型为结构体切片，有三个内容
	u := []Username{
		{Id: 1, Name: "张无忌", Age: 18},
		{Id: 2, Name: "周芷若", Age: 16},
		{Id: 3, Name: "谢逊", Age: 39},
	}
	//Execute负责渲染模板，并将b写入模板。
	_ = files.Execute(w, u)

	// w.Write([]byte("/index-struct"))
}
func main() {
	http.HandleFunc("/index", Index)
	http.HandleFunc("/index-map", IndexMap)
	http.HandleFunc("/index-struct", IndexStruct)
	// _ = http.ListenAndServe("", nil)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	} else {
		log.Print("ListenAndServe at 9090")
	}
}
