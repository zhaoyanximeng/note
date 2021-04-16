package main

import (
	"fmt"
	"sync"
	"time"
)

// 找出一个int切片里的所有偶数，再将他们乘以10

func Filter(list []int) chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for _, v := range list {
			if v%2 == 0 {
				c <- v
			}
		}
	}()

	return c
}

func Multiply(in chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			time.Sleep(time.Second * 1)
			out <- v * 10
		}
	}()

	return out
}

type Func func(list []int) chan int

type PipeFunc func(in chan int) chan int

// 管道函数
func PipelineFunction(args []int,c1 Func,cs ...PipeFunc) chan int {
	ret := c1(args)
	if len(cs) == 0 {
		return ret
	}

	result := make([]chan int,0)
	for index,c := range cs {
		if index == 0 {
			result = append(result,c(ret))
		} else {
			result = append(result,c(result[len(result) - 1]))
		}
	}

	return result[len(result) - 1]
}

// 多路复用
func NewPipelineFunction(args []int,c1 Func,cs ...PipeFunc) chan int {
	ret := c1(args)
	out := make(chan int)

	var wg sync.WaitGroup
	for _,c := range cs {
		getChan := c(ret)
		wg.Add(1)
		go func(input chan int) {
			defer wg.Done()
			for v := range input {
				out <- v
			}
		}(getChan)
	}

	go func() {
		defer close(out)
		wg.Wait()
	}()

	return out
}

func main() {
	nums := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}

	t := time.Now()

	ret := NewPipelineFunction(nums,Filter,Multiply,Multiply,Multiply,Multiply)

	for v := range ret {
		fmt.Println(v)
	}

	fmt.Println(time.Now().Sub(t))
}