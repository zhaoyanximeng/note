package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f func()
	var a *struct{}

	list := []interface{} {f,a}
	for _,v := range list {
		if reflect.ValueOf(v).IsNil() {
			fmt.Println(nil)
		}
	}
}
