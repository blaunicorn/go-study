package main

import (
	"fmt"
)

var C int = 2

// func ()  {

// }    // note: a function cannot be declared inside a function in Go language ,but anonymous functions can be placed.
//  note: Uppercase can be called by other packages, and lowercase is private to this package.
// func(input-parameter1 type,input-parameter2 type) (out-parameter1 type,out-parameter2 type) {
// 	function body
// }

func a1(data1 int, data2 string) {
	if data1 > 1 {
		fmt.Println(data1)
	} else {
		fmt.Println(data2)
	}
}
func a2(data1 int, data2 string) (ret1 int, ret2 string) {
	if data1 > 1 {
		fmt.Println(data1)
		return data1, "hahah"
	} else {
		fmt.Println(data2)
		return data1, "ok"
	}

}

// indefinite term parameter

func mo(data1 int, data2 ...string) {
	data2 = append(data2, "data1")
	for k, v := range data2 {
		fmt.Println(k, v)
	}
	// data2 = append(data2, strcov.Itoa(data1))
	fmt.Println(data1, data2)
}

// self executing function
// func(){
// 	fmt.Println("I am a self execution function!")
// }()

// closure function :  function returns a function
func closeFunction() func(num int) {
	return func(num int) {
		fmt.Println("this is a closure function", num)
	}
}

// deferred call function

func firstFunction() {
	fmt.Println("I want to be called in first!")
}
