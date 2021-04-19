package main

import "fmt"

type User struct {
	id int
}

type UserNew struct {
	id int
}

type MStruct struct {
	id int
	m map[string]string
}

func main() {
	a := User{id: 10}
	b := User{id: 10}
	// 打印true
	fmt.Println(a==b)

	//c := MStruct{id:10}
	//d := MStruct{id:10}
	// 无法比较，slice,map,func无法被比较
	//fmt.Println(c==d)

	c := User{id:10}
	d := UserNew{id:10}
	// 无法直接比较
	//fmt.Println(c == d)
	// 转化后可以比较
	fmt.Println(c == User(d))
}
