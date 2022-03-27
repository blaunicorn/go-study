package main

import (
	"05-init/lib1"
	"fmt"

	"05-init/lib2"
	_ "05-init/lib3" //别名可以匿名导入
)

func changeValue(p *int) {
	*p = 10
}
func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
	var a int = 1
	changeValue(&a)
	fmt.Println("a=", a)
}
