package main

import (
	"fmt"
	"time"
)

func Producer(out chan int) {
	defer close(out)
	for i := 0 ; i < 5 ; i ++ {
		out <- i * 2
		time.Sleep(time.Second * 2)
	}
}

func Consumer(out chan int) chan struct{} {
	r := make(chan struct{})
	go func() {
		defer func() {
			r <- struct{}{}
		}()

		for v := range out {
			fmt.Println(v)
		}
	}()

	return r
}

func main() {
	c := make(chan int)
	go Producer(c)
	r := Consumer(c)
	<- r
}


