package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int
	Price  int `json:"rmb"`
	Actors []string
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"xingye", "zhangzhi"}}
	//编码的过程 结构体---》json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error:", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码的过程 jsonStr--->struct
	myMovie := Movie{}
	err = json.Unmarshal(jsonStr, &myMovie)
	// err = json.Unmarshal([]byte(jsonStr), &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error:", err)
		return
	}
	fmt.Printf("%v\n", myMovie)
}
