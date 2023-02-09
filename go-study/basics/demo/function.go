package demo

// demo\function.go
import (
	"fmt"
	"strconv"
)

// Functional method

func FunctionalMethod() {

	// note: a function cannot be declared inside a function in Go language,
	// but anonymous functions can be placed.
	//  Uppercase can be called by other packages, and lowercase is private to this package.
	// func(input-parameter1 type,input-parameter2 type) (out-parameter1 type,out-parameter2 type) {
	// 	function body
	// }
	fmt.Println("Functional method:")
	a1(3, "4")
	a1(-1, "now it's -1")
	r1, r2 := a2(3, "now value <1")
	r3, r4 := a2(0, "now value <1")
	fmt.Println(r1, r2)
	fmt.Println(r3, r4)

	// Anonymous function
	b := func(data1 string) {
		fmt.Println(data1)
	}
	b("I am a anonmous function")

	a3(001, "1", "2", "4", "5")
	array1 := []string{"1", "2", "3", "4", "5"}
	a3(888, array1...)

	// self executing function
	func() {
		fmt.Println("I am a self execution function!")
	}()

	closeFunction()(4)

	defer firstFunction()     //  be defferred ,it will  carry out the final execution.
	fmt.Println("I am last?") // it  before defer funciton
}

// without out-parameter
func a1(data1 int, data2 string) {
	if data1 > 1 {
		fmt.Println(data1)
	} else {
		fmt.Println(data2)
	}
}

// with out-parameter
func a2(data1 int, data2 string) (ret1 int, ret2 string) {
	if data1 > 1 {
		fmt.Println(data1)
		return data1, "hahah"
	} else {
		fmt.Println(data2)
		return data1, "ok"
	}

}

// indefinite parameter
func a3(data1 int, data2 ...string) { // data2 is indefinite,it need place in last
	// data2 = append(data2, "data1")
	// for k, v := range data2 {
	// 	fmt.Println(k, v)
	// }
	str2 := strconv.Itoa(data1)
	data2 = append(data2, str2)
	fmt.Println(data1, data2)
}

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
