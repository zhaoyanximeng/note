package main

func main() {
	ch := make(chan int,3)
	ch <- 1
	ch <- 2
	ch <- 3

	// channel未关闭，for range会报错
	//for item := range ch {
	//	fmt.Println(item)
	//}

	// 关闭的channel读不会报错，会取到channel类型的零值
	//close(ch)
	//for {
	//	if item,ok := <- ch ; ok != false {
	//		fmt.Println(item)
	//	}
	//
	//}

	// 向已经关闭的channel写数据会报错
	close(ch)
	ch <- 4
}
