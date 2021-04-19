package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sync"
	"time"
)

var p *sync.Pool
type User struct {
	Name string
}

//func main()  {
//	p := sync.Pool{
//		New: func() interface{} {
//			log.Println("create user")
//			return &User{Name: "first"}
//		},
//	}
//	u1 := p.Get().(*User)
//	fmt.Println(u1)
//
//	u1.Name = "second"
//	p.Put(u1)
//
//	runtime.GC()
//	u2 := p.Get()
//	fmt.Println(u2)
//}

var retPool *sync.Pool

type Result struct {
	Message string `json:"message"`
	Status string `json:"status"`
	Logger interface{} `json:"-"`
}

func (this *Result) Success(ctx *gin.Context,msg string) {
	this.Status = "success"
	this.Message = msg
	ctx.JSON(200,this)
}

func (this *Result) Error(ctx *gin.Context,msg string) {
	this.Status = "success"
	this.Message = msg
	ctx.JSON(400,this)
}

func main() {
	retPool = &sync.Pool{
		New: func() interface{} {
			log.Println("created result")
			return &Result{}
		},
	}

	r := gin.New()
	r.Use(func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				ret := retPool.Get().(*Result)
				defer retPool.Put(ret)
				ret.Error(c,e.(string))
				c.Abort()
			}
		}()

		c.Next()
	})
	r.Handle("GET","/", func(c *gin.Context) {
		if time.Now().Unix() %2 == 0 {
			ret := retPool.Get().(*Result)
			defer retPool.Put(ret)
			ret.Success(c,"index")
		} else {

			panic("it's an error")
		}
	})

	r.Run(":8080")
}