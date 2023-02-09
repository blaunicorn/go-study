package main

import (
	"fmt"
	"sync"
	"time"
)

func SyncClass() {
	l := &sync.RWMutex{} // 异步互斥锁
	// l := &sync.Mutex{} // 异步互斥锁
	go lockFun(l)
	go lockFun(l)
	go lockFun(l)
	go readLockFun(l)
	go readLockFun(l)
	go readLockFun(l)
	for {
	}
	// lockFun()
	// lockFun()
	// lockFun()
	// lockFun()

}

func lockFun(lock *sync.RWMutex) {
	lock.Lock()
	fmt.Println("dododod")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}
func readLockFun(lock *sync.RWMutex) {
	lock.RLock()
	fmt.Println("hahahah")
	time.Sleep(1 * time.Second)
	lock.RUnlock()
}

// func lockFun() {
// 	fmt.Println("dododod")
// 	time.Sleep(1 * time.Second)
// }
