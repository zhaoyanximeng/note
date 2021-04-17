package main

import "fmt"

func main() {
	c := make(chan int,10)

	for i := 0 ; i < 10 ; i ++ {
		go func(input int) {
			c <- input
		}(i)
	}

	for i := 0 ; i < cap(c) ; i ++ {
		fmt.Println(i)
	}
}
