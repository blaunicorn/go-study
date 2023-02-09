package demo

import (
	"fmt"
	"sync"
)

func GoRun() {
	//  Run out of main program
	fmt.Println("goroutine & channel:")

	// way 1
	// go Run()
	// time.Sleep(1 * time.Second)

	// way 2
	// go Run()
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println(i)
	// }

	// true way：Program Manager
	var wg sync.WaitGroup
	wg.Add(1) // Set the number of cooperation programs to wait
	go Run1(&wg)
	wg.Wait() // Wait until all the cooperation programs are finished

	// chan channel is  communication bridge between goroutine
	// var chanName chan chanType
	// Five categories
	// c:=make(chan int)  //Readable and writable
	// 	var readChan <-chan int = c // only Readable
	// var setChan chan<-int = c  // only writable
	// c:=make(chan int,5)  //With cache
	// c:=make(chan int)  // No buffer

	// Define 2 chans
	chan1 := make(chan int, 5)

	chan1 <- 123456      // Data storage
	fmt.Println(<-chan1) // Fetching data

	// No buffer ,need goroutine channel
	chan2 := make(chan int)
	go func() {
		chan2 <- 222
	}()
	fmt.Println(<-chan2)

	fmt.Println("chan3 No buffer:")
	chan3 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			chan3 <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-chan3)
	}

	fmt.Println("chan4 With cache:")
	chan4 := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			chan4 <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-chan4)
	}

	fmt.Println("chan5 :")
	chan5 := make(chan int, 5)
	var readC <-chan int = chan5
	var writeC chan<- int = chan5
	writeC <- 5
	fmt.Println(<-readC)
	close(chan5) // close channel

	fmt.Println("chan6 :")
	chan6 := make(chan int, 5)
	chan6 <- 5
	chan6 <- 4
	chan6 <- 3
	chan6 <- 2
	chan6 <- 1
	close(chan6) // close channel
	for v := range chan6 {
		fmt.Println(v)
	}

	chan71 := make(chan int, 1)
	chan72 := make(chan int, 1)
	chan73 := make(chan int, 1)
	chan71 <- 71
	// chan72 <- 72
	// chan73 <- 73
	select {
	case m := <-chan71:
		fmt.Println("ch71:")
		fmt.Println(m)
	case <-chan72:
		fmt.Println("ch72:")
	case <-chan73:
		fmt.Println("ch73:")
	default:
		fmt.Println("no data")
	}

	fmt.Println("\n chan8 practical application:")
	chan8 := make(chan int)
	var readC8 <-chan int = chan8
	var writeC8 chan<- int = chan8
	go Setchan8(writeC8)
	Getchan8(readC8)
}

func Run() {
	fmt.Println("I ran.")
}

func Run1(wg *sync.WaitGroup) {
	fmt.Println("I ran in wg.")
	wg.Done() // //End of processing of a coroutine
}

func Setchan8(writeChan chan<- int) {
	for i := 0; i < 10; i++ {
		writeChan <- i * 3
	}
}

func Getchan8(readChan <-chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("i am from getFunciton，%d\n", <-readChan)
	}
}
