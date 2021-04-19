package main

import (
	"bytes"
	"fmt"
)

// 使用map实现一个简单的set，如果考虑线程安全用sync.Map

type EmptyStruct struct {}

type Set map[interface{}]EmptyStruct

func (this Set) Add(attrs ...interface{}) Set {
	for _,v := range attrs {
		this[v] = EmptyStruct{}
	}

	return this
}

func NewSet() Set {
	return make(map[interface{}]EmptyStruct)
}

func (this Set) String() string {
	var buf bytes.Buffer
	for k,_ := range this {
		if buf.Len() > 0 {
			buf.WriteString(",")
		}
		buf.WriteString(fmt.Sprintf("%v",k))
	}

	return buf.String()
}

func main() {
	set := NewSet().Add(1,2,3,4,5,3,4,"a")
	fmt.Println(set)
}
