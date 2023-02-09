package demo

import (
	"fmt"
	"sync"
	"time"
)

// demo\sync.go

func SyncClass() {
	// 锁的作用就是某个 协程 （线程）在访问某个资源时先锁住，防止其它协程的访问，等访问完毕解锁后其他协程再来加锁进行访问。

	// Go 语言 中的 sync 包提供了两种锁类型，分别为：sync.Mutex 和 sync.RWMutex，即互斥锁和 读写锁。
	// 	Mutex 是最简单的一种锁类型，同时也比较暴力，当一个 goroutine 获得了 Mutex 后，其他 goroutine 就只能乖乖等到这个 goroutine 释放该 Mutex。

	// RWMutex 相对友好些，是经典的单写多读模型。在读锁占用的情况下，会阻止写，但不阻止读，也就是多个 goroutine 可同时获取读锁（调用 RLock() 方法；而写锁（调用 Lock() 方法）会阻止任何其他 goroutine（无论读和写）进来，整个锁相当于由该 goroutine 独占。
	// 语法
	// 定义互斥锁变量 mutex
	// var mutex sync.Mutex
	// // 对需要访问的资源加锁
	// mutex.Lock()
	// // 资源访问结束解锁
	// mutex.Unlock()
	// mutex.Rlock()
	// mutex.Runlock()
	// Once.Do(FuncName)  无论被调用多少次，这里只执行一次
	// WaitGroup  Add(delta int) 设定需要Done的次数； Done() done一次； wait()在到达Done的次数时，一直阻塞
	// Map 并发字典 Store Load  LoadOrStore  Range Delete
	// Pool 并发池 Put  Get
	//  Cond  解锁通知
	fmt.Println("使用 sync.Mutex 互斥锁加锁操作")
	// 锁 需要指针形式
	// 互斥锁
	// lock := &sync.Mutex{}
	// go lockFun(lock)
	// go lockFun(lock)
	// go lockFun(lock)
	// // 无限循环
	// for {
	// }

	// 读写锁
	// lock := &sync.RWMutex{}
	// go lockFun1(lock)
	// go lockFun1(lock)
	// go lockFun1(lock)
	// go readLockFun(lock)
	// go readLockFun(lock)
	// go readLockFun(lock)
	// // 无限循环
	// for {
	// }

	// 只执行一次
	once := &sync.Once{}
	for i := 0; i < 10; i++ {
		once.Do(func() {
			fmt.Println(i)
		})
	}

	// 等待组
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		time.Sleep(5 * time.Second)
		wg.Done()
		fmt.Println("一个并行的函数")
	}()
	go func() {
		time.Sleep(6 * time.Second)
		wg.Done()
		fmt.Println("第二个个并行的函数")
	}()
	go func() {
		time.Sleep(5 * time.Second)
		wg.Done()
		fmt.Println("第三个并行的函数")
	}()
	wg.Wait()
	fmt.Println("三个并行都执行完毕")

	//  避免map异步读写错误，所以使用sync.map
	m := &sync.Map{}
	go func() {
		for {
			m.Store("london", 1)
		}
	}()
	go func() {
		count := 20
		for i := 0; i < count; i++ {
			fmt.Println(m.Load("london"))
		}
	}()
	time.Sleep(1000)
	m.Delete("london")
	m.LoadOrStore(3, 3)
	m.LoadAndDelete(3)

	m.Store(1, 11)
	m.Store(2, 22)
	m.Store(3, 33)
	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		time.Sleep(100 * time.Millisecond)
		return true
	})

	// 并发池
	p := &sync.Pool{}
	p.Put(1)
	p.Put(2)
	p.Put(3)
	p.Put(4)
	for i := 0; i < 4; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(p.Get())
	}

	cond := sync.NewCond(&sync.Mutex{})
	// cond.L.Lock()

	// cond.Wait()
	// cond.L.Unlock()

	// cond.Signal()    // 解开一个锁
	// cond.Broadcast() //广播解开所有的锁

	go func() {
		cond.L.Lock()
		fmt.Println("lock1")
		cond.Wait()
		fmt.Println("unlock1")
		cond.L.Unlock()
	}()
	go func() {
		cond.L.Lock()
		fmt.Println("lock2")
		cond.Wait()
		fmt.Println("unlock2")
		cond.L.Unlock()
	}()
	time.Sleep(2 * time.Second)
	// cond.Broadcast() //全通知
	cond.Signal() // 先通知一个
	time.Sleep(1 * time.Second)
	cond.Signal() //再通知要给
}

// 互斥锁
func lockFun(lock *sync.Mutex) {
	lock.Lock()
	fmt.Println("锁， 确保并发时，一次只执行一次")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}

func lockFun1(lock *sync.RWMutex) {
	lock.Lock()
	fmt.Println("锁， 写的时候，会排斥其他的读写")
	time.Sleep(1 * time.Second)
	lock.Unlock()
}
func readLockFun(lock *sync.RWMutex) {
	lock.RLock() // 在读取的时候，不会阻塞其他的读锁，但会排斥写锁。保证不能被串改
	fmt.Println("读锁")
	time.Sleep(1 * time.Second)
	lock.RUnlock()
}
