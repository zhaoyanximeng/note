package main

import (
	"fmt"
	"sync"
	"time"
)

// 通过有缓冲的channel控制
var pool chan struct{}

func job(index int) {
	time.Sleep(time.Millisecond * 500)
	fmt.Printf("执行完毕，序号：%d\n",index)
}

func main() {
	pool = make(chan struct{},10)
	wg := sync.WaitGroup{}

	for i := 0 ; i < 100 ; i ++ {
		pool <- struct{}{}
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			defer func() {
				<- pool
			}()
			job(index)
		}(i)
	}

	wg.Wait()
}
