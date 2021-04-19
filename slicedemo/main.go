package main

import "fmt"

// slice的浅拷贝和深拷贝

func main(){
	a := []int{1,2,3}

	// 浅拷贝,共享同一底层数组空间
	b := a
	b[1] = 100
	fmt.Println(a,b)

	// 深拷贝，不共享同一底层数组空间
	c := make([]int,len(a),cap(a))
	copy(c,a)
	c[1] = 1000
	fmt.Println(a,c)
}