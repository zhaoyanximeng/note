package main

import "fmt"

//func main() {
//	a := 1
//	// 打印2
//	//defer func() {
//	//	fmt.Println(a)
//	//}()
//
//	// 打印1
//	defer func(input int) {
//		fmt.Println(input)
//	}(a)
//
//	// 打印1
//	//defer fmt.Println(a)
//	a ++
//}

//func main() {
//	for i := 0 ; i < 3 ;i ++ {
//		// 保存了i的地址
//		// 打印出来结果是3
//		//defer func() {
//		//	fmt.Println(i)
//		//}()
//
//		// i值传递
//		// 打印出来结果是2 1 0
//		defer fmt.Println(i)
//	}
//}

func main() {
	defer func() {
		defer func() {
			fmt.Println(1)
		}()

		defer func() {
			fmt.Println(2)
		}()

		defer func() {
			fmt.Println(3)
		}()

		panic("触发异常1")
	}()

	panic("触发异常2")
}
