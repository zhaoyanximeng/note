package main

import (
	"fmt"
	"sync"
	"time"
)

var kv sync.Map

func Set(key string,value interface{},expire time.Duration) {
	kv.Store(key,value)
	time.AfterFunc(expire, func() {
		kv.Delete(key)
	})
}

func main() {
	Set("id",10,time.Second * 5)
	Set("name","ten",time.Second * 10)

	for {
		fmt.Println(kv.Load("id"))
		fmt.Println(kv.Load("name"))
		time.Sleep(time.Second)
	}
}
