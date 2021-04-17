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

// 打印结果是4 0 1 2 3
// 单核所有工作都是由一个线程执行（goroutine只是调度）
// GMP模型p会先执行最后一个创建的goroutine，再按顺序执行
//func main() {
//	runtime.GOMAXPROCS(1)
//	wg := sync.WaitGroup{}
//	for i := 0 ; i < 5 ;i ++ {
//		wg.Add(1)
//		go func(i int) {
//			defer wg.Done()
//			fmt.Println(i)
//		}(i)
//	}
//
//	wg.Wait()
//}