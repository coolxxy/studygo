package golang

import (
	"fmt"
	"sync"
	"time"
)

// channel 是 Go 语言线程间通信的一种机制，Go 从语言层面支持这种机制go.
// chan 分为 无缓冲 和 有缓冲两类
// chan 除了进行 goroutine 之间的消息传递之外，还可以实现一些线程管理的操作

// 无缓冲 chan 用于同步两个线程
// 通过一个 chan 保证正常打印
func Hello() {
	done := make(chan int)
	go func() {
		fmt.Println("Hello")
		done <- 1
	}()
	<-done
}

// 通过有缓冲 chan 实现并发数限制
func run(limitNum int, work []func()) {
	limit := make(chan int, limitNum)
	for _, w := range work {
		go func() {
			limit <- 1
			w()
			<-limit
		}()
	}
}

// 通过通道的关闭后可以接收一个零值的特性，实现线程管理
func worker(wg *sync.WaitGroup, exit chan bool) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-exit:
			return
		}
	}
}
func admin() {
	exit := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, exit)
	}

	time.Sleep(time.Second)
	close(exit)
	wg.Wait()
}
