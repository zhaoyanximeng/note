package main

import (
	"fmt"
	"math"
)

const abortIndex int8 = math.MaxInt8 / 2

// handlers为需要执行的handler链
// index是执行到的handler的索引
type Context struct {
	handlers []func(c *Context)
	index    int8
}

// 为Context的handler链添加handler
func (c *Context) Use(f func(c *Context)) {
	c.handlers = append(c.handlers, f)
}

func (c *Context) GET(path string, f func(c *Context)) {
	c.handlers = append(c.handlers, f)
}

// 更新Context里handler链的索引，方便下一个handler执行
func (c *Context) Next() {
	if c.index < int8(len(c.handlers)) {
		c.index++
		c.handlers[c.index](c)
	}
}

func (c *Context) Run() {
	c.handlers[0](c)
}

// 更新索引，使handler中断执行
func (c *Context) Abort() {
	c.index = abortIndex
}

func main() {
	c := &Context{}
	c.Use(Middleware1())
	c.Use(Middleware2())
	c.Use(Middleware3())
	c.GET("/", func(c *Context) {
		fmt.Println("get handler function")
	})
	c.Run()
}

func Middleware1() func(c *Context) {
	return func(c *Context) {
		fmt.Println("middleware1 start")
		c.Next()
		fmt.Println("middleware1 end")
	}
}

func Middleware2() func(c *Context) {
	return func(c *Context) {
		fmt.Println("middleware2 start")
		c.Abort()
		c.Next()
		fmt.Println("middleware2 end")
	}
}

func Middleware3() func(c *Context) {
	return func(c *Context) {
		fmt.Println("middleware3 start")
		c.Next()
		fmt.Println("middleware3 end")
	}
}
