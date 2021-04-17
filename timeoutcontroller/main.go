package main

import (
	"fmt"
	"time"
)

// 1.业务过程放到协程
// 2。把业务结果塞入channel
func job() chan string {
	ret := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ret <- "success"
	}()

	return ret
}

func run() (interface{},error) {
	c := job()
	select {
		case r := <- c:
			return r,nil
		case <- time.After(time.Second * 3):
			return nil,fmt.Errorf("time out")
	}
}

func main() {
	fmt.Println(run())
}
