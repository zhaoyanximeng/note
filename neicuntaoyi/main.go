package main

import "fmt"

// 内存逃逸：Go在编译阶段确定内存是分配在栈上还是堆上的一种行为
// 底层知识点：
// 1.栈内存分配和释放非常快
// 2.堆内存需要靠Go垃圾回收（占CPU）
// 3.通过逃逸分析，可以尽量把那些不需要分配到堆上的变量直接分配到栈上

// 命令：go build -gcflags=-m main.go

// 下面发生了内存逃逸
// 局部变量原本应该在栈中分配，在栈中回收，由于返回时被外部引用，发生了逃逸
//func test() []int {
//	a := []int{1,2,3}
//	a[1] = 4
//	return a  // 发生了逃逸
//}

// 下面发生了内存逃逸
// 参数为interface类型，比如fmt.Println(a ...interface{})
// 编译期间很难确定其参数的具体类型，也会产生逃逸
//func test() {
//	a := []int{1,2,3}
//	a[1] = 4
//	fmt.Println(a) // 发生了逃逸
//}
//
//func main() {
//	test()
//}

type User struct {
	id int
}

func NewUser() *User {
	return &User{id: 10}
}

func main()  {
	u := NewUser() // 这里发生了逃逸，u不传递到当前函数外部不建议使用指针，让NewUser()直接返回结构体本身即可
	fmt.Println(u)
}

