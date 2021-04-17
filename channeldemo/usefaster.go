package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 优胜劣模式（多个协程同时执行一个任务，取最先完成的）

func job() int {
	rand.Seed(time.Now().Unix())
	result := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(result))
	return result
}

func main() {
	c := make(chan int,5)
	for i := 0 ;i < 5 ;i ++ {
		go func() {
			c <- job()
		}()
	}

	fmt.Println("最快用了",<-c,"秒")
}